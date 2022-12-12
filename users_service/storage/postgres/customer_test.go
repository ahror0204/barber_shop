package postgres

import (
	"testing"

	pb "github.com/barber_shop/users_service/genproto"
	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/require"
)

func createCustomer(t *testing.T) *pb.Customer {
	customer, err := repoCustomer.CreateCustomer(&pb.Customer{
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

	return customer
}

func deleteCustomer(id string, t *testing.T) {
	err := repoCustomer.DeleteCustomer(&pb.ID{Id: id})
	require.NoError(t, err)
}

func TestCreateCustomer(t *testing.T) {
	customer := createCustomer(t)
	deleteCustomer(customer.Id, t)
}

func TestUpdateCustomer(t *testing.T) {
	cust := createCustomer(t)

	customer, err := repoCustomer.UpdateCustomer(&pb.Customer{
		Id:          cust.Id,
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
	deleteCustomer(customer.Id, t)
}

func TestGetCustomerByID(t *testing.T) {
	cust := createCustomer(t)

	customer, err := repoCustomer.GetCustomerByID(&pb.ID{Id: cust.Id})

	require.NoError(t, err)
	require.NotEmpty(t, customer)
	deleteCustomer(customer.Id, t)
}

func TestGetListCustomers(t *testing.T) {
	customer := createCustomer(t)
	customers, err := repoCustomer.GetListCustomers(&pb.GetCustomerParams{
		Limit: 10,
		Page:  1,
	})

	require.NoError(t, err)
	require.GreaterOrEqual(t, len(customers.Customers), 1)
	deleteCustomer(customer.Id, t)
}

func TestGetCustomerByEmail(t *testing.T) {
	cust := createCustomer(t)

	customer, err := repoCustomer.GetCustomerByEmail(&pb.Email{Email: cust.Email})
	
	require.NoError(t, err)
	require.NotEmpty(t, customer)
	deleteCustomer(customer.Id, t)
}

func TestUpdateCustomerPassword(t *testing.T) {
	customer := createCustomer(t)

	err := repoCustomer.UpdateCustomerPassword(&pb.UpdateCustomerPasswordRequest{ID: customer.Id, Password: faker.Password()})
	require.NoError(t, err)
	deleteCustomer(customer.Id, t)
}