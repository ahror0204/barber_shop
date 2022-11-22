package postgres

import (
	"testing"

	pb "github.com/barber_shop/user_service/genproto"
	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/require"
)

func createUser(t *testing.T) *pb.ID {
	id, err := repoUser.CreateUser(&pb.User{
		FirstName:   faker.FirstName(),
		LastName:    faker.LastName(),
		PhoneNumber: faker.Phonenumber(),
		Email:       faker.Email(),
		UserName:    faker.Username(),
		Password:    faker.Password(),
		Gender:      "male",
		ImageUrl:    faker.URL(),
	})

	require.NoError(t, err)
	require.NotEmpty(t, id)

	return id
}

func deleteUser(id *pb.ID, t *testing.T) {
	err := repoUser.DeleteUser(id)
	require.NoError(t, err)
}

func TestCreateUser(t *testing.T) {
	id := createUser(t)
	deleteUser(id, t)
}

func TestUpdateUser(t *testing.T) {
	id := createUser(t)

	user, err := repoUser.UpdateUser(&pb.User{
		Id:          id.Id,
		FirstName:   faker.FirstName(),
		LastName:    faker.LastName(),
		PhoneNumber: faker.Phonenumber(),
		Email:       faker.Email(),
		UserName:    faker.Username(),
		Password:    faker.Password(),
		Gender:      "male",
		ImageUrl:    faker.URL(),
	})

	require.NoError(t, err)
	require.NotEmpty(t, user)
	deleteUser(id, t)
}

func TestGetUserByID(t *testing.T) {
	id := createUser(t)

	user, err := repoUser.GetUserByID(id)

	require.NoError(t, err)
	require.NotEmpty(t, user)
	deleteUser(id, t)
}

func TestGetAllUsers(t *testing.T) {
	id := createUser(t)
	users, err := repoUser.GetAllUsers(&pb.GetUserParams{
		Limit: 10,
		Page:  1,
	})

	require.NoError(t, err)
	require.GreaterOrEqual(t, len(users.Users), 1)
	deleteUser(id, t)
}
