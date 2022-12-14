package postgres

import (
	"log"
	"os"
	"testing"

	"github.com/barber_shop/users_service/config"
	"github.com/barber_shop/users_service/pkg/db"
	"github.com/barber_shop/users_service/pkg/logger"
	"github.com/barber_shop/users_service/storage/repo"
)

var (
	repoCustomer repo.CustomerStorageI
	repoSalon    repo.SalonStorageI
	repoStaff    repo.StaffStorageI
)

func TestMain(m *testing.M) {
	cfg := config.Load("./../..")

	db, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}

	repoCustomer = NewCustomerRepo(db)
	repoSalon = NewSalonRepo(db)
	repoStaff = NewStaffRepo(db)
	os.Exit(m.Run())
}
