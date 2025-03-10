package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	Init()
	assert.NotNil(t, Log)
	Log.Info("Logger initialized")
	Sync()
}
