package postgres

import (
	"database/sql"
	"fmt"
	"time"

	pbu "github.com/barber_shop/users_service/genproto/users_service"
	"github.com/gofrs/uuid"

	"github.com/barber_shop/users_service/storage/repo"
	"github.com/jmoiron/sqlx"
)

type staffRepo struct {
	db *sqlx.DB
}

func NewStaffRepo(db *sqlx.DB) repo.StaffStorageI {
	return &staffRepo{db}
}

func (s *staffRepo) CreateStaff(staff *pbu.Staff) (*pbu.Staff, error) {
	query := `INSERT INTO staff(
		id,
		salon_id,
		first_name,
		last_name,
		phone_number,
		email,
		user_name,
		password,
		type,
		image_url
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
	RETURNING id, created_at`
	ID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	err = s.db.QueryRow(query,
		ID,
		staff.SalonId,
		staff.FirstName,
		staff.LastName,
		staff.PhoneNumber,
		staff.Email,
		staff.UserName,
		staff.Password,
		staff.Type,
		staff.ImageUrl,
	).Scan(&staff.Id, &staff.CreatedAt)

	if err != nil {
		return nil, err
	}
	return staff, nil
}

func (u *staffRepo) UpdateStaff(staff *pbu.Staff) (*pbu.Staff, error) {
	var rstaff pbu.Staff
	query := `UPDATE staff SET
	first_name = $1,
	last_name = $2,
	phone_number = $3,
	email = $4,
	user_name = $5,
	password = $6,
	image_url = $7,
	updated_at = $8
	WHERE id = $9
	RETURNING id, salon_id, first_name, last_name, phone_number, 
	email, user_name, password, type, image_url, created_at, updated_at`

	updateAT := time.Now()

	err := u.db.QueryRow(query,
		staff.FirstName,
		staff.LastName,
		staff.PhoneNumber,
		staff.Email,
		staff.UserName,
		staff.Password,
		staff.ImageUrl,
		updateAT,
		staff.Id,
	).Scan(
		&rstaff.Id,
		&rstaff.SalonId,
		&rstaff.FirstName,
		&rstaff.LastName,
		&rstaff.PhoneNumber,
		&rstaff.Email,
		&rstaff.UserName,
		&rstaff.Password,
		&rstaff.Type,
		&rstaff.ImageUrl,
		&rstaff.CreatedAt,
		&rstaff.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &rstaff, nil
}

func (u *staffRepo) GetStaffByID(ID *pbu.ID) (*pbu.Staff, error) {
	var (
		updateAT sql.NullTime
		rstaff   pbu.Staff
	)
	query := `SELECT
		id,
		salon_id,
		first_name,
		last_name,
		phone_number,
		email,
		user_name,
		password,
		type,
		image_url,
		created_at,
		updated_at
	FROM staff
	WHERE id = $1 AND deleted_at IS NULL`

	err := u.db.QueryRow(query, ID.Id).Scan(
		&rstaff.Id,
		&rstaff.SalonId,
		&rstaff.FirstName,
		&rstaff.LastName,
		&rstaff.PhoneNumber,
		&rstaff.Email,
		&rstaff.UserName,
		&rstaff.Password,
		&rstaff.Type,
		&rstaff.ImageUrl,
		&rstaff.CreatedAt,
		&updateAT,
	)
	if err != nil {
		return nil, err
	}

	if updateAT.Valid {
		rstaff.UpdatedAt = updateAT.Time.String()
	}

	return &rstaff, nil
}

func (u *staffRepo) GetListStaff(params *pbu.GetListParams) (*pbu.ListStaff, error) {
	var (
		staff []*pbu.Staff
		count int64
	)
	offset := (params.Page - 1) * params.Limit
	limit := fmt.Sprintf("LIMIT %d OFFSET %d", params.Limit, offset)
	filter := ""
	if params.Search != "" {
		str := "%" + params.Search + "%"
		filter = fmt.Sprintf(`
			WHERE first_name ILIKE '%s' OR last_name ILIKE '%s' OR email ILIKE '%s' 
			OR phone_number ILIKE '%s' AND deleted_at IS NULL
		`, str, str, str, str)
	}else{filter = " WHERE deleted_at IS NULL "}
	query := `SELECT
		id,
		salon_id,
		first_name,
		last_name,
		phone_number,
		email,
		user_name,
		password,
		type,
		image_url,
		created_at,
		updated_at
	FROM staff ` + filter + `ORDER BY created_at DESC ` + limit

	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			updateAT sql.NullTime
			st       pbu.Staff
		)
		err = rows.Scan(
			&st.Id,
			&st.SalonId,
			&st.FirstName,
			&st.LastName,
			&st.PhoneNumber,
			&st.Email,
			&st.UserName,
			&st.Password,
			&st.Type,
			&st.ImageUrl,
			&st.CreatedAt,
			&updateAT,
		)

		if err != nil {
			return nil, err
		}

		if updateAT.Valid {
			st.UpdatedAt = time.Time.String(updateAT.Time)
		}

		staff = append(staff, &st)
	}

	countQuery := `SELECT count(1) FROM staff` + filter
	err = u.db.QueryRow(countQuery).Scan(&count)
	if err != nil {
		return nil, err
	}

	return &pbu.ListStaff{
		Staff: staff,
		Count: count,
	}, nil
}

func (s *staffRepo) DeleteStaff(ID *pbu.ID) error {
	deletedAT := time.Now()
	query := `UPDATE staff SET deleted_at=$1 WHERE id = $2`

	result, err := s.db.Exec(query, deletedAT, ID.Id)
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


func (c *staffRepo) GetStaffByEmail(email *pbu.Email) (*pbu.Staff, error) {
	var (
		updateAT  sql.NullTime
		rStaff pbu.Staff
	)

	query := `SELECT
			id,
			salon_id,
			first_name,
			last_name,
			phone_number,
			email,
			user_name,
			password,
			type,
			image_url,
			created_at,
			updated_at
		FROM staff
		WHERE email = $1 AND deleted_at IS NULL`

	err := c.db.QueryRow(query, email.Email).Scan(
		&rStaff.Id,
		&rStaff.SalonId,
		&rStaff.FirstName,
		&rStaff.LastName,
		&rStaff.PhoneNumber,
		&rStaff.Email,
		&rStaff.UserName,
		&rStaff.Password,
		&rStaff.Type,
		&rStaff.ImageUrl,
		&rStaff.CreatedAt,
		&updateAT,
	)
	if err != nil {
		return nil, err
	}

	if updateAT.Valid {
		rStaff.UpdatedAt = updateAT.Time.String()
	}
	return &rStaff, nil
}


func (s *staffRepo) UpdateStaffPassword(req *pbu.UpdatePasswordRequest) error {
	query := `UPDATE staff SET password=$1 WHERE id=$2`
	_, err := s.db.Exec(query, req.Password, req.ID)
	return err
}