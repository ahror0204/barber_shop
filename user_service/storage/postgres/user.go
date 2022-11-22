package postgres

import (
	"database/sql"
	"fmt"
	"time"

	pb "github.com/barber_shop/user_service/genproto"
	"github.com/barber_shop/user_service/storage/repo"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) repo.UserStorageI {
	return &userRepo{db}
}

func (u *userRepo) CreateUser(user *pb.User) (*pb.ID, error) {
	var id string
	query := `INSERT INTO users(
		id,
		first_name,
		last_name,
		phone_number,
		email,
		user_name,
		gender,
		password,
		image_url
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)
	RETURNING id`
	ID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	err = u.db.QueryRow(
		query,
		ID,
		user.FirstName,
		user.LastName,
		user.PhoneNumber,
		user.Email,
		user.UserName,
		user.Gender,
		user.Password,
		user.ImageUrl,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &pb.ID{Id: id}, nil
}

func (u *userRepo) UpdateUser(user *pb.User) (*pb.User, error) {
	var ruser pb.User
	query := `UPDATE users SET
		first_name=$1,
		last_name=$2,
		phone_number=$3,
		email=$4,
		user_name=$5,
		gender=$6,
		image_url=$7,
		updated_at=$8
	WHERE id = $9
	RETURNING id, first_name, last_name, phone_number, 
	email, user_name, gender, image_url, created_at, updated_at`

	updateAT := time.Now()

	err := u.db.QueryRow(
		query,
		user.FirstName,
		user.LastName,
		user.PhoneNumber,
		user.Email,
		user.UserName,
		user.Gender,
		user.ImageUrl,
		updateAT,
		user.Id,
	).Scan(
		&ruser.Id,
		&ruser.FirstName,
		&ruser.LastName,
		&ruser.PhoneNumber,
		&ruser.Email,
		&ruser.UserName,
		&ruser.Gender,
		&ruser.ImageUrl,
		&ruser.CreatedAt,
		&ruser.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &ruser, nil
}

func (u *userRepo) GetUserByID(ID *pb.ID) (*pb.User, error) {
	var updateAT sql.NullTime
	var ruser pb.User
	query := `SELECT
		id,
		first_name,
		last_name,
		phone_number,
		email,
		user_name,
		gender,
		image_url,
		created_at,
		updated_at
	FROM users
	WHERE id = $1 AND deleted_at IS NULL`

	err := u.db.QueryRow(query, ID.Id).Scan(
		&ruser.Id,
		&ruser.FirstName,
		&ruser.LastName,
		&ruser.PhoneNumber,
		&ruser.Email,
		&ruser.UserName,
		&ruser.Gender,
		&ruser.ImageUrl,
		&ruser.CreatedAt,
		&updateAT,
	)
	if err != nil {
		return nil, err
	}

	if updateAT.Valid {
		ruser.UpdatedAt = updateAT.Time.String()
	}

	return &ruser, nil
}

func (u *userRepo) GetAllUsers(params *pb.GetUserParams) (*pb.AllUsers, error) {
	var (
		users []*pb.User
		count int64
	)
	offset := (params.Page - 1) * params.Limit
	limit := fmt.Sprintf("LIMIT %d OFFSET %d", params.Limit, offset)
	filter := ""
	if params.Search != "" {
		str := "%" + params.Search + "%"
		filter = fmt.Sprintf(`
			WHERE first_name ILIKE '%s' OR last_name ILIKE '%s' OR email ILIKE '%s' 
			OR username ILIKE '%s' OR phone_number ILIKE '%s'
		`, str, str, str, str, str)
	}
	query := `SELECT
		id,
		first_name,
		last_name,
		phone_number,
		email,
		user_name,
		gender,
		image_url,
		created_at,
		updated_at
	FROM users ` + filter + `ORDER BY created_at DESC ` + limit

	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		var(
			updateAT sql.NullTime
			user pb.User
		)
		err = rows.Scan(
			&user.Id,
			&user.FirstName,
			&user.LastName,
			&user.PhoneNumber,
			&user.Email,
			&user.UserName,
			&user.Gender,
			&user.ImageUrl,
			&user.CreatedAt,
			&updateAT,
		)

		if err != nil {
			return nil, err
		}

		if updateAT.Valid {
			user.UpdatedAt = time.Time.String(updateAT.Time)
		}

		users = append(users, &user)
	}

	countQuery := `SELECT count(1) FROM users` + filter
	err = u.db.QueryRow(countQuery).Scan(&count)
	if err != nil {
		return nil, err
	}

	return &pb.AllUsers{
		Users: users,
		Count: count,
	}, nil
}

func (u *userRepo) DeleteUser(ID *pb.ID) error {
	deletedAT := time.Now()
	query := `UPDATE users SET deleted_at=$1 WHERE id = $2`

	result, err := u.db.Exec(query, deletedAT, ID.Id)
	if err != nil {
		return err
	}

	rowsCount, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsCount == 0 {
		return sql.ErrNoRows
	}

	return nil
}
