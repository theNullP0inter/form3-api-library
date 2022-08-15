package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigLive(t *testing.T) {
	cfg := MockConfigLive

	assert.Equal(t, cfg.GetBaseUrl(), LiveEndPoint)
	assert.NotEqual(t, cfg.GetBaseUrl(), TestEndPoint)
}

func TestConfigTest(t *testing.T) {
	cfg := MockConfigTest

	assert.Equal(t, cfg.GetBaseUrl(), TestEndPoint)
	assert.NotEqual(t, cfg.GetBaseUrl(), LiveEndPoint)
}
