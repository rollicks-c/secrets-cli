package bitwarden

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/rollicks-c/secretblendproviders/bitwarden"
	"github.com/rollicks-c/secrets-cli/internal/config"
	"strings"
)

func connect() (*bitwarden.Client, error) {

	conf := config.Profiles().LoadCurrent().Data

	client, err := bitwarden.NewClient(conf.Backends.Bitwarden.DataDir)
	if err != nil {
		return nil, err
	}

	if err := unlockVault(client); err != nil {
		return nil, err
	}

	if err := client.Check(); err != nil {
		return nil, err
	}

	return client, nil
}

func unlockVault(c *bitwarden.Client) error {

	locked, err := c.IsLocked()
	if err != nil {
		return err
	}
	if !locked {
		return nil
	}

	password := ""
	for {
		password, err = promptConfString("enter master password:", "")
		if err != nil {
			return err
		}
		if password == "" {
			continue
		}
		break
	}

	if err := c.Unlock(password); err != nil {
		return err
	}
	return nil
}
func promptConfString(prompt, defaultValue string) (string, error) {

	dataPrompt := promptui.Prompt{
		Label:   fmt.Sprintf("%s (%s)", prompt, defaultValue),
		Default: defaultValue,
		Mask:    '*',
		Validate: func(s string) error {
			if strings.TrimSpace(s) == "" {
				return fmt.Errorf("value required")
			}
			return nil
		},
	}
	value, err := dataPrompt.Run()
	if err != nil {
		return "", err
	}
	return value, nil

}
