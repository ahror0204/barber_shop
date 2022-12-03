package postgres

import (
	"database/sql"
	"fmt"
	"time"

	pb "github.com/barber_shop/users_service/genproto"
	"github.com/barber_shop/users_service/storage/repo"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

type customerRepo struct {
	db *sqlx.DB
}

func NewCustomer(db *sqlx.DB) repo.CustomerStorageI {
	return &customerRepo{db}
}

func (u *customerRepo) CreateCustomer(customer *pb.Customer) (*pb.ID, error) {
	var id string
	query := `INSERT INTO customers(
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
		customer.FirstName,
		customer.LastName,
		customer.PhoneNumber,
		customer.Email,
		customer.UserName,
		customer.Gender,
		customer.Password,
		customer.ImageUrl,
	).Scan(&id)

	if err != nil {
		return nil, err
	}

	return &pb.ID{Id: id}, nil
}

func (u *customerRepo) UpdateCustomer(customer *pb.Customer) (*pb.Customer, error) {
	var rCustomer pb.Customer
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
	email, user_name, gender, image_url, created_at, updated_at`

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
		&rCustomer.ImageUrl,
		&rCustomer.CreatedAt,
		&rCustomer.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &rCustomer, nil
}

func (u *customerRepo) GetCustomerByID(ID *pb.ID) (*pb.Customer, error) {
	var (
		updateAT sql.NullTime
	 	rCustomer pb.Customer
	)
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

func (u *customerRepo) GetListCustomers(params *pb.GetCustomerParams) (*pb.AllCustomers, error) {
	var (
		customers []*pb.Customer
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
	FROM customers ` + filter + `ORDER BY created_at DESC ` + limit

	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			updateAT sql.NullTime
			customer pb.Customer
		)
		err = rows.Scan(
			&customer.Id,
			&customer.FirstName,
			&customer.LastName,
			&customer.PhoneNumber,
			&customer.Email,
			&customer.UserName,
			&customer.Gender,
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

	return &pb.AllCustomers{
		Customers: customers,
		Count:     count,
	}, nil
}

func (u *customerRepo) DeleteCustomer(ID *pb.ID) error {
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

func (c *customerRepo) GetCustomerByEmail(email *pb.Email) (*pb.Customer, error)  {
	var (
		updateAT sql.NullTime
	 	rCustomer pb.Customer
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