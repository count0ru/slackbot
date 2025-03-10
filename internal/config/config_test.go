package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	cfg, err := LoadConfig("tests/config.test.yaml")
	assert.NoError(t, err)

	assert.Equal(t, 8080, cfg.General.Port)
	assert.Equal(t, 9090, cfg.Metrics.Port)
	assert.Equal(t, "test.db", cfg.Database.Path)
	assert.Equal(t, "https://test-domain.atlassian.net/rest/api/3/issue", cfg.Tracker.URL)
	assert.Equal(t, "TEST", cfg.Ticket.Project.Key)
}
