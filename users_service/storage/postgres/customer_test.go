package postgres

import (
	"testing"

	pbu "github.com/barber_shop/users_service/genproto/users_service"
	"github.com/bxcodec/faker/v4"
	"github.com/stretchr/testify/require"
)

func createCustomer(t *testing.T) *pbu.Customer {
	customer, err := repoCustomer.CreateCustomer(&pbu.Customer{
		FirstName:   faker.FirstName(),
		LastName:    faker.LastName(),
		PhoneNumber: faker.Phonenumber(),
		Email:       faker.Email(),
		UserName:    faker.Username(),
		Password:    faker.Password(),
		Gender:      "male",
		Type:        "user",
		ImageUrl:    faker.URL(),
	})

	require.NoError(t, err)
	require.NotEmpty(t, customer)

	return customer
}

func deleteCustomer(id string, t *testing.T) {
	err := repoCustomer.DeleteCustomer(&pbu.ID{Id: id})
	require.NoError(t, err)
}

func TestCreateCustomer(t *testing.T) {
	customer := createCustomer(t)
	deleteCustomer(customer.Id, t)
}

func TestUpdateCustomer(t *testing.T) {
	cust := createCustomer(t)

	customer, err := repoCustomer.UpdateCustomer(&pbu.Customer{
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

	customer, err := repoCustomer.GetCustomerByID(&pbu.ID{Id: cust.Id})

	require.NoError(t, err)
	require.NotEmpty(t, customer)
	deleteCustomer(customer.Id, t)
}

func TestGetListCustomers(t *testing.T) {
	customer := createCustomer(t)
	customers, err := repoCustomer.GetListCustomers(&pbu.GetListParams{
		Limit: 10,
		Page:  1,
	})

	require.NoError(t, err)
	require.GreaterOrEqual(t, len(customers.Customers), 1)
	deleteCustomer(customer.Id, t)
}

func TestGetCustomerByEmail(t *testing.T) {
	cust := createCustomer(t)

	customer, err := repoCustomer.GetCustomerByEmail(&pbu.Email{Email: cust.Email})

	require.NoError(t, err)
	require.NotEmpty(t, customer)
	deleteCustomer(customer.Id, t)
}

func TestUpdateCustomerPassword(t *testing.T) {
	customer := createCustomer(t)

	err := repoCustomer.UpdateCustomerPassword(&pbu.UpdatePasswordRequest{ID: customer.Id, Password: faker.Password()})
	require.NoError(t, err)
	deleteCustomer(customer.Id, t)
}
