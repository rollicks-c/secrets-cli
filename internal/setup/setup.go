package setup

import (
	"github.com/rollicks-c/secretblendproviders/envvar"
)

func EnableEnvVarSupport() error {
	if err := envvar.RegisterGlobally(); err != nil {
		return err
	}
	return nil
}
