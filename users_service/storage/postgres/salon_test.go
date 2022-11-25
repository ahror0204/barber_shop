package postgres

import (
	"testing"

	pb "github.com/barber_shop/users_service/genproto"
	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/require"
)

func createSalon(t *testing.T) *pb.ID {

	id, err := repoSalon.CreateSalon(&pb.Salon{
		Name:        faker.Name(),
		PhoneNumber: faker.Phonenumber(),
		Email:       faker.Email(),
		Rating:      45,
		Address:     faker.MacAddress(),
		Latitude:    faker.LATITUDE,
		Longitude:   faker.LONGITUDE,
		StartTime:   faker.Timestamp(),
		EndTime:     faker.Timestamp(),
		ImageUrl:    faker.URL(),
	})

	require.NoError(t, err)
	require.NotEmpty(t, id)

	return id
}

func deleteSalon(id *pb.ID, t *testing.T) {
	err := repoSalon.DeleteSalon(id)
	require.NoError(t, err)
}

func TestCreateSalon(t *testing.T) {
	id := createSalon(t)
	deleteSalon(id, t)
}

func TestUpdateSalon(t *testing.T) {
	id := createSalon(t)
	salon, err := repoSalon.UpdateSalon(&pb.Salon{
		Id:          id.Id,
		Name:        faker.Name(),
		PhoneNumber: faker.Phonenumber(),
		Email:       faker.Email(),
		Rating:      4,
		Address:     faker.MacAddress(),
		Latitude:    faker.LATITUDE,
		Longitude:   faker.LONGITUDE,
		StartTime:   faker.Timestamp(),
		EndTime:     faker.Timestamp(),
		ImageUrl:    faker.URL(),
	})

	require.NoError(t, err)
	require.NotEmpty(t, salon)

	deleteSalon(id, t)
}

func TestGetSalonByID(t *testing.T) {
	id := createSalon(t)
	salon, err := repoSalon.GetSalonByID(id)

	require.NoError(t, err)
	require.NotEmpty(t, salon)

	deleteSalon(id, t)
}

func TestGetListSalons(t *testing.T) {
	id := createSalon(t)
	salons, err := repoSalon.GetListSalons(&pb.GetSalonsParams{
		Page:  1,
		Limit: 10,
	})

	require.NoError(t, err)
	require.GreaterOrEqual(t, len(salons.Salons), 1)

	deleteSalon(id, t)
}
