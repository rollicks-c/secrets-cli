package bitwarden

import (
	"fmt"
	"github.com/rollicks-c/secrets-cli/internal/client/commands/params"
	"github.com/rollicks-c/term"
	"github.com/urfave/cli/v2"
	"strings"
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
	itemID := pc.GetString(0)
	keyExp := pc.GetString(1)
	copyToClipboard := FlagCopyToClipboard.Get(c)
	key, ok := fuzzySearchKey(keyExp)
	if !ok {
		return fmt.Errorf("key [%s] not found in item [%s]", keyExp, itemID)
	}

	// switch case
	if isTOTP(key) {
		return getTOTP(itemID, copyToClipboard)
	} else {
		return getSecret(itemID, key, copyToClipboard)
	}

}

func Find(c *cli.Context) error {

	// collect params
	var pc = params.NewCollector(c.Args().Slice())
	if err := pc.Validate(1); err != nil {
		return err
	}
	exp := strings.Join(pc.GetAll(), " ")

	// search in bitwarden
	client, err := connect()
	if err != nil {
		return err
	}
	if err := client.Sync(); err != nil {
		return err
	}
	item, err := client.Find(exp)
	if err != nil {
		return err
	}

	// print result
	if len(item) == 0 {
		term.Warnf("No items found for [%s]\n", exp)
		return nil
	}
	term.Successf("Found %d items:\n", len(item))
	for _, i := range item {
		term.Successf("  %s\t%s\n", i.Id, i.Name)
	}
	return nil

}
