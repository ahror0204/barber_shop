package main

import (
	"fmt"
	"log"

	"github.com/barber_shop/users/api/v1"
	"github.com/barber_shop/users/config"
	"github.com/barber_shop/users/storage"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load("./")
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", 
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DataBase,
	)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("failed to open connection %v:", err)
	}
	st := storage.NewStoragePg(db)
	server := api.NewService(st)
	err = server.Run(cfg.HTTPPort)
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}


	

}
