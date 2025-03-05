package commands

import (
	"github.com/rollicks-c/secrets-cli/internal/config"
	"github.com/urfave/cli/v2"
	"os/exec"
)

type OpenCommand struct {
}

func (ac OpenCommand) Run(c *cli.Context) error {

	// create URLS
	conf := config.Profiles().LoadCurrent().Data
	vtURL := conf.Backends.Vault.Address

	// open
	cmd := exec.Command("brave-browser", vtURL)
	if err := cmd.Start(); err != nil {
		return err
	}

	return nil
}
