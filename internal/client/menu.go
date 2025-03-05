package client

import (
	"github.com/rollicks-c/secrets-cli/internal/client/commands"
	"github.com/rollicks-c/secrets-cli/internal/client/commands/alias"
	"github.com/rollicks-c/secrets-cli/internal/client/commands/bitwarden"
	"github.com/rollicks-c/secrets-cli/internal/client/commands/vault"
	"github.com/urfave/cli/v2"
)

const (
	cmdVault     = "vault"
	cmdBitwarden = "bitwarden"
	cmdFind      = "find"

	cmdAliasRun    = "alias-run"
	cmdAliasManage = "alias-manage"

	cmdProfile = "profile"
	cmdOpen    = "open"
)

func createCommands() []*cli.Command {

	return []*cli.Command{
		createAliasRunCommand(),
		createAliasManageCommand(),

		createVaultCommand(),
		createBitwardenCommand(),
		createFindCommand(),

		createProfileCommand(),
		createOpenCommand(),
	}
}

func createAliasRunCommand() *cli.Command {
	return &cli.Command{
		Name:            cmdAliasRun,
		Aliases:         []string{"ar", "a", "r", "run"},
		Action:          alias.Run,
		HideHelpCommand: true,
	}
}

func createAliasManageCommand() *cli.Command {
	return &cli.Command{
		Name:    cmdAliasManage,
		Aliases: []string{"am"},
		Action:  alias.List,

		Subcommands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "add {ALIAS} {VAULT PATH}",
				Action:  alias.Add,
			},
			{
				Name:    "remove",
				Aliases: []string{"rm"},
				Usage:   "remove {ALIAS}",
				Action:  alias.Remove,
			},
		},
	}
}

func createProfileCommand() *cli.Command {
	return &cli.Command{
		Name:            cmdProfile,
		Aliases:         []string{"p"},
		Action:          commands.SwitchProfileCommand{}.Run,
		HideHelpCommand: true,
		Subcommands: []*cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "list",
				Action:  commands.ListProfileCommand{}.Run,
			},
			{
				Name:            "switch",
				Aliases:         []string{"s"},
				Usage:           "switch",
				Action:          commands.SwitchProfileCommand{}.Run,
				HideHelpCommand: true,
			},
		},
	}
}

func createVaultCommand() *cli.Command {
	return &cli.Command{
		Name:    cmdVault,
		Aliases: []string{"vt"},
		Usage:   "{PATH} {KEY}",
		Action:  vault.Get,
		Flags:   []cli.Flag{vault.FlagCopyToClipboard},
	}
}

func createBitwardenCommand() *cli.Command {
	return &cli.Command{
		Name:    cmdBitwarden,
		Aliases: []string{"bw"},
		Usage:   "{ITEM} {KEY}",
		Action:  bitwarden.Get,
		Flags:   []cli.Flag{bitwarden.FlagCopyToClipboard},
	}
}

func createFindCommand() *cli.Command {
	return &cli.Command{
		Name:    cmdFind,
		Aliases: []string{"f"},
		Usage:   "{EXP}",
		Action:  bitwarden.Find,
	}
}

func createOpenCommand() *cli.Command {
	return &cli.Command{
		Name:   cmdOpen,
		Action: commands.OpenCommand{}.Run,
	}
}
