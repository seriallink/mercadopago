package mp

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPaymentMethods(t *testing.T) {
	pms, err := client.GetPaymentMethods()
	assert.NoError(t, err)
	assert.NotZero(t, len(pms))
	spew.Dump(pms)
}
