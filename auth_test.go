package mp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const clientId = "6897814360277806"
const clientSecret = "nQeEuaPIVWiOBHXHGyPbjsgYdka9aPv3"
const publicKey = "TEST-fa412681-b4c4-4b2e-b6f5-23a716673006"
const accessToken = "TEST-6897814360277806-042515-a9c8081578f60060c13e0e4ec3ab358b__LD_LA__-253816740"

var client *Client

func init() {
	client = NewClient(clientId, clientSecret, publicKey, accessToken)
}

func TestAuthorize(t *testing.T) {
	err := client.Authorize()
	assert.NoError(t, err)
	assert.NotEmpty(t, client.AuthToken.AccessToken, "AccessToken cannot be empty")
	assert.NotEmpty(t, client.AuthToken.RefreshToken, "RefreshToken cannot be empty")
}
