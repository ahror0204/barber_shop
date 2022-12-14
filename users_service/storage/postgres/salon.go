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

type salonRepo struct {
	db *sqlx.DB
}

func NewSalonRepo(db *sqlx.DB) repo.SalonStorageI {
	return &salonRepo{db}
}

func (s *salonRepo) CreateSalon(salon *pbu.Salon) (*pbu.Salon, error) {
	query := `INSERT INTO salon(
		id,
		name,
		phone_number,
		email,
		rating,
		address,
		latitude,
		longitude,
		start_time,
		end_time,
		image_url)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
	RETURNING id, created_at
	`
	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	
	err = s.db.QueryRow(query,
		id,
		salon.Name,
		salon.PhoneNumber,
		salon.Email,
		salon.Rating,
		salon.Address,
		salon.Latitude,
		salon.Longitude,
		salon.StartTime,
		salon.EndTime,
		salon.ImageUrl,
	).Scan(&salon.Id, &salon.CreatedAt)
	if err != nil {
		return nil, err
	}
	return salon, nil
}

func (s *salonRepo) UpdateSalon(salon *pbu.Salon) (*pbu.Salon, error) {
	query := `UPDATE salon SET
		name=$1,
		phone_number=$2,
		email=$3,
		rating=$4,
		address=$5,
		latitude=$6,
		longitude=$7,
		start_time=$8,
		end_time=$9,
		image_url=$10,
		updated_at=$11
	WHERE id = $12
	RETURNING created_at, updated_at
	`
	updatedAT := time.Now()
	err := s.db.QueryRow(query,
		salon.Name,
		salon.PhoneNumber,
		salon.Email,
		salon.Rating,
		salon.Address,
		salon.Latitude,
		salon.Longitude,
		salon.StartTime,
		salon.EndTime,
		salon.ImageUrl,
		updatedAT,
		salon.Id,
	).Scan(&salon.CreatedAt, &salon.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return salon, nil
}

func (s *salonRepo) GetSalonByID(ID *pbu.ID) (*pbu.Salon, error) {
	var (
		salon     pbu.Salon
		updatedAT sql.NullTime
	)
	query := `SELECT 
		id,
		name,
		phone_number,
		email,
		rating,
		address,
		latitude,
		longitude,
		start_time,
		end_time,
		image_url,
		created_at,
		updated_at
	FROM salon
	WHERE id=$1
	`

	err := s.db.QueryRow(query, ID.Id).Scan(
		&salon.Id,
		&salon.Name,
		&salon.PhoneNumber,
		&salon.Email,
		&salon.Rating,
		&salon.Address,
		&salon.Latitude,
		&salon.Longitude,
		&salon.StartTime,
		&salon.EndTime,
		&salon.ImageUrl,
		&salon.CreatedAt,
		&updatedAT,
	)
	if err != nil {
		return nil, err
	}
	if updatedAT.Valid {
		salon.UpdatedAt = updatedAT.Time.String()
	}
	return &salon, nil
}

func (s *salonRepo) GetListSalons(params *pbu.GetListParams) (*pbu.AllSalons, error) {
	var (
		salons    []*pbu.Salon
		count     int64
		updatedAT sql.NullTime
	)
	offset := (params.Page - 1) * params.Limit
	limit := fmt.Sprintf("LIMIT %d OFFSET %d", params.Limit, offset)
	filter := ""
	if params.Search != "" {
		str := "%" + params.Search + "%"
		filter = fmt.Sprintf(`
			WHERE name ILIKE '%s' OR rating ILIKE '%s' OR email ILIKE '%s' 
			OR phone_number ILIKE '%s' AND deleted_at IS NULL
		`, str, str, str, str)
	}
	query := `SELECT 
		id,
		name,
		phone_number,
		email,
		rating,
		address,
		latitude,
		longitude,
		start_time,
		end_time,
		image_url,
		created_at,
		updated_at
	FROM salon
	` + filter + ` ORDER BY created_at DESC ` + limit
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var salon pbu.Salon
		err := rows.Scan(
			&salon.Id,
			&salon.Name,
			&salon.PhoneNumber,
			&salon.Email,
			&salon.Rating,
			&salon.Address,
			&salon.Latitude,
			&salon.Longitude,
			&salon.StartTime,
			&salon.EndTime,
			&salon.ImageUrl,
			&salon.CreatedAt,
			&updatedAT,
		)
		if err != nil {
			return nil, err
		}
		if updatedAT.Valid {
			salon.UpdatedAt = updatedAT.Time.String()
		}

		salons = append(salons, &salon)
	}

	countQuery := `SELECT count(1) FROM salon` + filter
	err = s.db.QueryRow(countQuery).Scan(&count)
	if err != nil {
		return nil, err
	}
	return &pbu.AllSalons{
		Salons: salons,
		Count:  count,
	}, nil
}

func (s *salonRepo) DeleteSalon(ID *pbu.ID) error {
	deletedAT := time.Now()
	query := "UPDATE salon SET deleted_at = $1 WHERE id = $2"
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
