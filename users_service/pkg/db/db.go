package db

import (
	"fmt"
	"log"

	"github.com/barber_shop/users_service/config"
	"github.com/barber_shop/users_service/pkg/logger"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectToDB(cfg config.Config) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DataBase,
	)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal("failed to open connection:", logger.Error(err))
	}
	return db, nil
}
