package mp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateCardToken(t *testing.T) {
	cardtoken := &CardToken{
		PublicKey:       client.PublicKey,
		ExpirationYear:  2020,
		ExpirationMonth: 10,
		FirstSixDigits:  "123456",
		CardHolder: &CardHolder{
			Name: "Marcelo Monaco",
			Identification: &Identification{
				Number: "000000000000000000000000000000000000000000000000000000000000000000000000",
			},
		},
	}
	err := client.CreateCardToken(cardtoken)
	assert.NoError(t, err)
	assert.NotEmpty(t, cardtoken.CardTokenId, "CardTokenId cannot be empty")
}

func TestGetCardToken(t *testing.T) {
	cardtoken := &CardToken{
		PublicKey:   client.PublicKey,
		CardTokenId: "42da2523de96a648d634c7d24ad8eaec",
	}
	err := client.GetCardToken(cardtoken)
	assert.NoError(t, err)
	assert.Equal(t, "0000000000000000", cardtoken.CardHolder.Identification.Number)
}
