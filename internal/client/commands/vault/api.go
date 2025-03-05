package vault

import (
	"fmt"
	"github.com/rollicks-c/secrets-cli/internal/client/commands/params"
	"github.com/rollicks-c/secrets-cli/internal/config"
	"github.com/rollicks-c/secrets-cli/internal/logic/ui"
	vaultui "github.com/rollicks-c/vaultclient/ui"
	"github.com/rollicks-c/vaultclient/vault"
	"github.com/urfave/cli/v2"
	"path"
)

var (
	FlagCopyToClipboard = &cli.BoolFlag{
		Name:    "copy",
		Value:   true,
		Aliases: []string{"c"},
	}
)

func Get(c *cli.Context) error {

	// collect params
	var pc = params.NewCollector(c.Args().Slice())
	if err := pc.Validate(2); err != nil {
		return err
	}
	vtPath := pc.GetString(0)
	vtKey := pc.GetString(1)
	copyToClipboard := FlagCopyToClipboard.Get(c)

	// create client
	conf := config.Profiles().LoadCurrent().Data
	vaultAddr := conf.Backends.Vault.Address
	vaultToken, _ := config.LoadVaultToken(conf)
	prompter := createPersistenceMiddleware(conf, vaultui.CreateVaultTokenPrompter(vaultAddr))
	vt, err := vault.NewClient(conf.Backends.Vault.Address, vault.WithTokenPrompt(vaultToken, prompter))
	if err != nil {
		return err
	}

	// lookup
	vtSecret, err := vt.LoadSecret(vtPath)
	if err != nil {
		return err
	}
	if vtSecret == nil {
		return fmt.Errorf("secret [%s] not	found", vtPath)
	}
	vtKey, vtValue, ok := vtSecret.GetItemFuzzy(vtKey)
	if !ok {
		return fmt.Errorf("key [%s] not found in secret [%s]", vtKey, vtPath)
	}

	// provide
	if copyToClipboard {
		if err := ui.PutInClipboard(vtValue); err != nil {
			return err
		}
		notification := fmt.Sprintf("VT secret [%s] of [%s] ready in clipboard", vtKey, path.Base(vtPath))
		if err := ui.NotifyUser(notification); err != nil {
			return err
		}
	} else {
		fmt.Print(vtValue)
	}

	return nil
}

func createPersistenceMiddleware(conf config.Configuration, prompter func() (string, error)) func() (string, error) {
	return func() (string, error) {

		token, err := prompter()
		if err != nil {
			return "", err
		}

		config.SaveVaultToken(conf, token)
		return token, nil
	}
}
