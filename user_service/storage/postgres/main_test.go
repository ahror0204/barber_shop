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

var repoI repo.UserStorageI

func TestMain(m *testing.M) {
	cfg := config.Load("./../..")

	db, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	repoI = NewUser(db)
	os.Exit(m.Run())
}
