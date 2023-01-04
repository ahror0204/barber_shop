package postgres

import (
	"database/sql"
	"fmt"
	"time"

	pbu "github.com/barber_shop/users_service/genproto/users_service"
	"github.com/barber_shop/users_service/storage/repo"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

type customerRepo struct {
	db *sqlx.DB
}

func NewCustomerRepo(db *sqlx.DB) repo.CustomerStorageI {
	return &customerRepo{db}
}

func (u *customerRepo) CreateCustomer(customer *pbu.Customer) (*pbu.Customer, error) {
	query := `INSERT INTO customers(
		id,
		first_name,
		last_name,
		phone_number,
		email,
		user_name,
		gender,
		password,
		type,
		image_url
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
	RETURNING id, created_at`
	ID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	err = u.db.QueryRow(query,
		ID,
		customer.FirstName,
		customer.LastName,
		customer.PhoneNumber,
		customer.Email,
		customer.UserName,
		customer.Gender,
		customer.Password,
		customer.Type,
		customer.ImageUrl,
	).Scan(&customer.Id, &customer.CreatedAt)

	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (u *customerRepo) UpdateCustomer(customer *pbu.Customer) (*pbu.Customer, error) {
	var rCustomer pbu.Customer
	query := `UPDATE customers SET
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
	email, user_name, gender, type, image_url, created_at, updated_at`

	updateAT := time.Now()

	err := u.db.QueryRow(
		query,
		customer.FirstName,
		customer.LastName,
		customer.PhoneNumber,
		customer.Email,
		customer.UserName,
		customer.Gender,
		customer.ImageUrl,
		updateAT,
		customer.Id,
	).Scan(
		&rCustomer.Id,
		&rCustomer.FirstName,
		&rCustomer.LastName,
		&rCustomer.PhoneNumber,
		&rCustomer.Email,
		&rCustomer.UserName,
		&rCustomer.Gender,
		&rCustomer.Type,
		&rCustomer.ImageUrl,
		&rCustomer.CreatedAt,
		&rCustomer.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &rCustomer, nil
}

func (u *customerRepo) GetCustomerByID(ID *pbu.ID) (*pbu.Customer, error) {
	var (
		updateAT  sql.NullTime
		rCustomer pbu.Customer
	)
	query := `SELECT
		id,
		first_name,
		last_name,
		phone_number,
		email,
		user_name,
		gender,
		type,
		image_url,
		created_at,
		updated_at
	FROM customers
	WHERE id = $1 AND deleted_at IS NULL`

	err := u.db.QueryRow(query, ID.Id).Scan(
		&rCustomer.Id,
		&rCustomer.FirstName,
		&rCustomer.LastName,
		&rCustomer.PhoneNumber,
		&rCustomer.Email,
		&rCustomer.UserName,
		&rCustomer.Gender,
		&rCustomer.Type,
		&rCustomer.ImageUrl,
		&rCustomer.CreatedAt,
		&updateAT,
	)
	if err != nil {
		return nil, err
	}

	if updateAT.Valid {
		rCustomer.UpdatedAt = updateAT.Time.String()
	}

	return &rCustomer, nil
}

func (u *customerRepo) GetListCustomers(params *pbu.GetListParams) (*pbu.AllCustomers, error) {
	var (
		customers []*pbu.Customer
		count     int64
	)
	offset := (params.Page - 1) * params.Limit
	limit := fmt.Sprintf("LIMIT %d OFFSET %d", params.Limit, offset)
	filter := ""
	if params.Search != "" {
		str := "%" + params.Search + "%"
		filter = fmt.Sprintf(`
			WHERE first_name ILIKE '%s' OR last_name ILIKE '%s' OR email ILIKE '%s' 
			OR user_name ILIKE '%s' OR phone_number ILIKE '%s' AND deleted_at IS NULL
		`, str, str, str, str, str)
	} else {
		filter = " WHERE deleted_at IS NULL "
	}
	query := `SELECT
		id,
		first_name,
		last_name,
		phone_number,
		email,
		user_name,
		gender,
		type,
		image_url,
		created_at,
		updated_at
	FROM customers ` + filter + `ORDER BY created_at DESC ` + limit

	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			updateAT sql.NullTime
			customer pbu.Customer
		)
		err = rows.Scan(
			&customer.Id,
			&customer.FirstName,
			&customer.LastName,
			&customer.PhoneNumber,
			&customer.Email,
			&customer.UserName,
			&customer.Gender,
			&customer.Type,
			&customer.ImageUrl,
			&customer.CreatedAt,
			&updateAT,
		)

		if err != nil {
			return nil, err
		}

		if updateAT.Valid {
			customer.UpdatedAt = time.Time.String(updateAT.Time)
		}

		customers = append(customers, &customer)
	}

	countQuery := `SELECT count(1) FROM customers` + filter
	err = u.db.QueryRow(countQuery).Scan(&count)
	if err != nil {
		return nil, err
	}

	return &pbu.AllCustomers{
		Customers: customers,
		Count:     count,
	}, nil
}

func (u *customerRepo) DeleteCustomer(ID *pbu.ID) error {
	deletedAT := time.Now()
	query := `UPDATE customers SET deleted_at=$1 WHERE id = $2`

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

func (c *customerRepo) GetCustomerByEmail(email *pbu.Email) (*pbu.Customer, error) {
	var (
		updateAT  sql.NullTime
		rCustomer pbu.Customer
	)

	query := `SELECT
			id,
			first_name,
			last_name,
			phone_number,
			email,
			user_name,
			password,
			gender,
			type,
			image_url,
			created_at,
			updated_at
		FROM customers
		WHERE email = $1 AND deleted_at IS NULL`

	err := c.db.QueryRow(query, email.Email).Scan(
		&rCustomer.Id,
		&rCustomer.FirstName,
		&rCustomer.LastName,
		&rCustomer.PhoneNumber,
		&rCustomer.Email,
		&rCustomer.UserName,
		&rCustomer.Password,
		&rCustomer.Gender,
		&rCustomer.Type,
		&rCustomer.ImageUrl,
		&rCustomer.CreatedAt,
		&updateAT,
	)
	if err != nil {
		return nil, err
	}

	if updateAT.Valid {
		rCustomer.UpdatedAt = updateAT.Time.String()
	}
	return &rCustomer, nil
}

func (c *customerRepo) UpdateCustomerPassword(req *pbu.UpdatePasswordRequest) error {
	query := `UPDATE customers SET password=$1 WHERE id=$2`
	_, err := c.db.Exec(query, req.Password, req.ID)
	return err
}
