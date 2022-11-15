package postgres

import (
	"github.com/barber_shop/users/storage/repo"
	str "github.com/barber_shop/users/structures"
	"github.com/gofrs/uuid"
	"github.com/jmoiron/sqlx"
)

type userRepo struct {
	db *sqlx.DB
}

func NewUser(db *sqlx.DB) repo.UserStorageI{
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
		gender,
		password,
		image_url
	) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
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
		user.Gender,
		user.Password,
		user.ImageURL,
	).Scan(&id)

	if err != nil {
		return "", nil
	}

	return id, nil
}

