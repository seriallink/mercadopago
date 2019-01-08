package mp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCustomer(t *testing.T) {
	customer := &Customer{
		CustomerId: "693516-UkMF8LxlzTVVMY",
	}
	err := client.GetCustomer(customer)
	assert.NoError(t, err)
	assert.NotEmpty(t, customer.CustomerId)
	assert.Equal(t, "693516-UkMF8LxlzTVVMY", customer.CustomerId)
	assert.Equal(t, "teste@teste.com", customer.Email)
}

func TestSearchCustomer(t *testing.T) {
	customer := &Customer{
		Email: "teste@teste.com",
	}
	err := client.SearchCustomer(customer)
	assert.NoError(t, err)
	assert.NotEmpty(t, customer.CustomerId)
	assert.Equal(t, "693516-UkMF8LxlzTVVMY", customer.CustomerId)
}

func TestCreateCustomer(t *testing.T) {
	customer := &Customer{
		Email: "abc@teste.com",
	}
	err := client.CreateCustomer(customer)
	assert.NoError(t, err)
	assert.NotEmpty(t, customer.CustomerId)
	assert.Equal(t, "abc@teste.com", customer.Email)
}

func TestSaveCustomer(t *testing.T) {
	customer := &Customer{
		CustomerId: "693516-UkMF8LxlzTVVMY",
		Email:      "outro@email.com",
		FirstName:  "Teste",
		LastName:   "Testando",
	}
	err := client.SaveCustomer(customer)
	assert.NoError(t, err)
	assert.NotEmpty(t, customer.CustomerId)
	assert.Equal(t, "teste@teste.com", customer.Email)
	assert.Equal(t, "Teste", customer.FirstName)
}

func TestRemoveCustomer(t *testing.T) {
	customer := &Customer{
		CustomerId: "229777231-9lp4UMV8d5ViRE",
	}
	err := client.RemoveCustomer(customer)
	assert.NoError(t, err)
	assert.Equal(t, "abc@teste.com", customer.Email)
}

func TestGetCustomerCards(t *testing.T) {
	customer := &Customer{
		CustomerId: "693516-UkMF8LxlzTVVMY",
	}
	err := client.GetCustomer(customer)
	assert.NoError(t, err)
	assert.NotEmpty(t, customer.CustomerId)
	assert.Equal(t, "693516-UkMF8LxlzTVVMY", customer.CustomerId)
	assert.Equal(t, "teste@teste.com", customer.Email)
}

func TestNewCustomerCard(t *testing.T) {
	customer := &Customer{
		CustomerId: "693516-UkMF8LxlzTVVMY",
		Cards: []Card{{
			CustomerId:      "693516-UkMF8LxlzTVVMY",
			FirstSixDigits:  "000000",
			LastFourDigits:  "0000",
			ExpirationMonth: 10,
			ExpirationYear:  2020,
			SecurityCode: &SecurityCode{
				Length: 3,
			},
		}},
	}

	err := client.NewCustomerCard(customer)
	assert.NoError(t, err)
	assert.NotEmpty(t, customer.Cards[0].CardId)
	assert.Equal(t, "693516-UkMF8LxlzTVVMY", customer.CustomerId)
}
