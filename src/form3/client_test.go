package form3

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/theNullP0inter/form3-api-library/src/common"
)

func TestForm3Client(t *testing.T) {
	cli := NewForm3Client(common.MockConfigTest)

	assert.IsType(t, &Form3Client{}, cli)

	assert.NotNil(t, cli.Account)
	assert.NotNil(t, cli.client)

}
