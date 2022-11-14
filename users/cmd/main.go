package main

import (
	"fmt"
	"github.com/barber_shop/users/config"

	_ "github.com/lib/pq"
)

func main(){
	cfg := config.Load("./..")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DataBase,
	)

	

}