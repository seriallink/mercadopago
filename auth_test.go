package mp

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

const clientId = ""
const clientSecret = ""
const accessToken = ""
const sandbox = true

func TestNewClient(t *testing.T) {
	client := NewClient(clientId, clientSecret, "", sandbox)
	err := client.Authorize()
	assert.NoError(t, err)
	assert.NotEmpty(t, client.AuthToken.AccessToken, "AccessToken cannot be empty")
	assert.NotEmpty(t, client.AuthToken.RefreshToken, "RefreshToken cannot be empty")
}
