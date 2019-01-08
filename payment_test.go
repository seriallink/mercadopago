package mp

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCreateTicketPayment(t *testing.T) {
	payment := &Payment{
		DateOfExpiration:  NewIso8601(time.Now().Add(time.Hour * 48)).Pointer(),
		TransactionAmount: 10,
		Description:       "Payment Test",
		PaymentMethodId:   "bolbradesco",
		Payer: &Payer{
			Email:     "test@test.com",
			FirstName: "John",
			LastName:  "Doe",
			Identification: &Identification{
				Type:   "CPF",
				Number: "000000000-00",
			},
			Address: &Address{
				StreetName:   "Major Basilio",
				StreetNumber: "661A",
				Neighborhood: "Vila Bertioga",
				City:         "Sao Paulo",
				State:        "SP",
				ZipCode:      "03181-010",
			},
		},
	}

	err := client.CreatePayment(payment)
	assert.NoError(t, err)
	assert.NotZero(t, payment.Id)

}
