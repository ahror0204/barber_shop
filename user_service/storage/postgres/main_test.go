package postgres

import (
	"log"
	"os"
	"testing"

	"github.com/barber_shop/user_service/config"
	"github.com/barber_shop/user_service/pkg/db"
	"github.com/barber_shop/user_service/pkg/logger"
	"github.com/barber_shop/user_service/storage/repo"
)

var (
	repoUser repo.UserStorageI
	repoSalon repo.SalonStorageI
)

func TestMain(m *testing.M) {
	cfg := config.Load("./../..")

	db, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	repoUser = NewUser(db)
	repoSalon = NewSalonRepo(db)
	os.Exit(m.Run())
}
