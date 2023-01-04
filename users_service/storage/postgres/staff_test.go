package postgres

import (
	"testing"

	pbu "github.com/barber_shop/users_service/genproto/users_service"
	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/require"
)

func createStaff(t *testing.T) *pbu.Staff {
	salon := createSalon(t)
	staff, err := repoStaff.CreateStaff(&pbu.Staff{
		SalonId:     salon.Id,
		FirstName:   faker.FirstName(),
		LastName:    faker.LastName(),
		PhoneNumber: faker.Phonenumber(),
		Email:       faker.Email(),
		UserName:    faker.Username(),
		Password:    faker.Password(),
		Type:        "user",
		ImageUrl:    faker.URL(),
	})

	require.NoError(t, err)
	require.NotEmpty(t, staff)

	return staff
}

func deleteStaff(id string, t *testing.T) {
	err := repoStaff.DeleteStaff(&pbu.ID{Id: id})
	require.NoError(t, err)
}

func TestCreateStaff(t *testing.T) {
	staff := createStaff(t)
	deleteStaff(staff.Id, t)
}

func TestUpdateStaff(t *testing.T) {
	staff := createStaff(t)
	res, err := repoStaff.UpdateStaff(staff)

	require.NoError(t, err)
	require.NotEmpty(t, res)
	deleteStaff(staff.Id, t)
}

func TestGetStaffByID(t *testing.T) {
	staff := createStaff(t)
	s, err := repoStaff.GetStaffByID(&pbu.ID{Id: staff.Id})

	require.NoError(t, err)
	require.NotEmpty(t, s)
	deleteStaff(staff.Id, t)
}

func TestGetListStaff(t *testing.T) {
	staff := createStaff(t)
	s, err := repoStaff.GetListStaff(&pbu.GetListParams{
		Limit: 10,
		Page:  1,
	})

	require.NoError(t, err)
	require.GreaterOrEqual(t, len(s.Staff), 1)
	deleteStaff(staff.Id, t)
}

func TestGetStaffByEmail(t *testing.T) {
	s := createStaff(t)

	staff, err := repoStaff.GetStaffByEmail(&pbu.Email{Email: s.Email})

	require.NoError(t, err)
	require.NotEmpty(t, staff)
	deleteStaff(staff.Id, t)
}

func TestUpdateStaffPassword(t *testing.T) {
	staff := createStaff(t)

	err := repoStaff.UpdateStaffPassword(&pbu.UpdatePasswordRequest{ID: staff.Id, Password: faker.Password()})
	require.NoError(t, err)
	deleteStaff(staff.Id, t)
}
