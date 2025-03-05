package main

import (
	"github.com/rollicks-c/secrets-cli/internal/client"
	"github.com/rollicks-c/secrets-cli/internal/setup"
	"github.com/rollicks-c/term"
	"os"
)

func main() {

	if err := setup.EnableEnvVarSupport(); err != nil {
		term.Failf("%s\n", err.Error())
		return
	}

	if err := client.CreateClient().Run(os.Args); err != nil {
		term.Failf("%s\n", err.Error())
	}

}
