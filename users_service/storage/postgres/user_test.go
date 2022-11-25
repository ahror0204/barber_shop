package postgres

import (
	"testing"

	pb "github.com/barber_shop/users_service/genproto"
	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/require"
)

func createCustomer(t *testing.T) *pb.ID {
	id, err := repoCustomer.CreateCustomer(&pb.Customer{
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

func deleteCustomer(id *pb.ID, t *testing.T) {
	err := repoCustomer.DeleteCustomer(id)
	require.NoError(t, err)
}

func TestCreateCustomer(t *testing.T) {
	id := createCustomer(t)
	deleteCustomer(id, t)
}

func TestUpdateCustomer(t *testing.T) {
	id := createCustomer(t)

	customer, err := repoCustomer.UpdateCustomer(&pb.Customer{
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
	require.NotEmpty(t, customer)
	deleteCustomer(id, t)
}

func TestGetCustomerByID(t *testing.T) {
	id := createCustomer(t)

	customer, err := repoCustomer.GetCustomerByID(id)

	require.NoError(t, err)
	require.NotEmpty(t, customer)
	deleteCustomer(id, t)
}

func TestGetListCustomers(t *testing.T) {
	id := createCustomer(t)
	customers, err := repoCustomer.GetListCustomers(&pb.GetCustomerParams{
		Limit: 10,
		Page:  1,
	})

	require.NoError(t, err)
	require.GreaterOrEqual(t, len(customers.Customers), 1)
	deleteCustomer(id, t)
}
