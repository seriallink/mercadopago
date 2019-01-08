package mp

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetIdentificationTypes(t *testing.T) {
	ids, err := client.GetIdentificationTypes()
	assert.NoError(t, err)
	assert.NotZero(t, len(ids))
	spew.Dump(ids)
}
