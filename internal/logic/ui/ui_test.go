package ui

import (
	"github.com/gen2brain/beeep"
	"github.com/rollicks-c/secrets-cli/internal/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNotifications(t *testing.T) {

	{
		err := beeep.Notify(config.AppLabel, "test", "assets/warning.png")
		assert.NoError(t, err)
	}
	{
		err := beeep.Notify(config.AppLabel, "test", "")
		assert.NoError(t, err)
	}

}
