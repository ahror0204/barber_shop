package postgres

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/barber_shop/users/storage/repo"
	str "github.com/barber_shop/users/structures"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) repo.UserStorageI {
	return &userRepo{db: db}
}

func (u *userRepo) CreateUser(user *str.User) (string, error) {
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
		return "", err
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
		user.ImageURL,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (u *userRepo) UpdateUser(user *str.User) (*str.User, error) {
	var ruser str.User
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
		user.ImageURL,
		updateAT,
		user.ID,
	).Scan(
		&ruser.ID,
		&ruser.FirstName,
		&ruser.LastName,
		&ruser.PhoneNumber,
		&ruser.Email,
		&ruser.UserName,
		&ruser.Gender,
		&ruser.ImageURL,
		&ruser.CreatedAT,
		&ruser.UpdatedAT,
	)
	if err != nil {
		return nil, err
	}

	return &ruser, nil
}

func (u *userRepo) GetUserByID(ID string) (*str.User, error) {
	var updateAT sql.NullTime
	var ruser str.User
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

	err := u.db.QueryRow(query, ID).Scan(
		&ruser.ID,
		&ruser.FirstName,
		&ruser.LastName,
		&ruser.PhoneNumber,
		&ruser.Email,
		&ruser.UserName,
		&ruser.Gender,
		&ruser.ImageURL,
		&ruser.CreatedAT,
		&updateAT,
	)
	if err != nil {
		return nil, err
	}

	if updateAT.Valid {
		ruser.UpdatedAT = &updateAT.Time
	}

	return &ruser, nil
}

func (u *userRepo) GetAllUsers(params *str.GetUsersParams) (*str.AllUsers, error) {
	var (
		users []str.User
		count int64
	)
	offset := (params.Page-1)*params.Limit
	limit := fmt.Sprintf("LIMIT %d OFFSET %d", params.Limit, offset)
	filter := ""
	if params.Search != "" {
		str := "%"+params.Search+"%"
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
	FROM users ` + filter + `
	ORDER BY created_at DESC ` + limit
	
	err := u.db.Select(&users, query)
	
	if err != nil {
		return nil, err
	}

	countQuery := `SELECT count(1) FROM users`+filter
	err = u.db.QueryRow(countQuery).Scan(&count)
	if err != nil {
		return nil, err
	}

	return &str.AllUsers{
		Users: users,
		Count: count,
	}, nil
}

func (u *userRepo) DeleteUser(ID string) error {
	deleteAT := time.Now()
	query := `UPDATE users SET
		deleted_at=$1
	WHERE id = $2`

	_, err := u.db.Exec(query, deleteAT, ID)
	if err != nil {
		return err
	}

	return nil
}