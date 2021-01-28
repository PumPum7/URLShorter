package utils

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestUtils(t *testing.T) {
	assert.NotNil(t, Cfg)
	assert.NotNil(t, Log)
	assert.FileExists(t, Cfg.Logging.Logfile)
}

func TestLoggerLevel(t *testing.T) {
	level := Cfg.Logging.LogLevel
	assert.Equal(t, strings.ToLower(level), Log.Level.String())
}
