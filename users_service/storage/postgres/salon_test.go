package postgres

import (
	"testing"

	pb "github.com/barber_shop/users_service/genproto"
	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/require"
)

func createSalon(t *testing.T) *pb.Salon {

	salon, err := repoSalon.CreateSalon(&pb.Salon{
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
	require.NotEmpty(t, salon)

	return salon
}

func deleteSalon(id string, t *testing.T) {
	err := repoSalon.DeleteSalon(&pb.ID{Id: id})
	require.NoError(t, err)
}

func TestCreateSalon(t *testing.T) {
	salon := createSalon(t)
	deleteSalon(salon.Id, t)
}

func TestUpdateSalon(t *testing.T) {
	s := createSalon(t)
	salon, err := repoSalon.UpdateSalon(&pb.Salon{
		Id:          s.Id,
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

	deleteSalon(s.Id, t)
}

func TestGetSalonByID(t *testing.T) {
	s := createSalon(t)
	salon, err := repoSalon.GetSalonByID(&pb.ID{Id: s.Id})

	require.NoError(t, err)
	require.NotEmpty(t, salon)

	deleteSalon(s.Id, t)
}

func TestGetListSalons(t *testing.T) {
	s := createSalon(t)
	salons, err := repoSalon.GetListSalons(&pb.GetSalonsParams{
		Page:  1,
		Limit: 10,
	})

	require.NoError(t, err)
	require.GreaterOrEqual(t, len(salons.Salons), 1)

	deleteSalon(s.Id, t)
}
