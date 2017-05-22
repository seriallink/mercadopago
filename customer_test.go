package mp

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewCustomer(t *testing.T) {
	client := NewClient(clientId, clientSecret, accessToken, sandbox)
	customer := &Customer{
		Email: "teste@teste.com",
	}
	err := client.NewCustomer(customer)
	assert.NoError(t,err)
	assert.NotEmpty(t, customer.CustomerId)
	assert.Equal(t, "teste@teste.com", customer.Email)
}