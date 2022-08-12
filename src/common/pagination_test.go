package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPagination(t *testing.T) {
	p := NewPagination(1, 10)

	assert.IsType(t, Pagination{}, p)
	assert.Equal(t, p.Number, 1)
	assert.Equal(t, p.Size, 10)
}
