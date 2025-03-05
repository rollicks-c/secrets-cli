package commands

import (
	"github.com/rollicks-c/secrets-cli/internal/config"
	"github.com/rollicks-c/term"
	"github.com/urfave/cli/v2"
	"strings"
)

type SwitchProfileCommand struct {
}

func (d SwitchProfileCommand) Run(c *cli.Context) error {

	if c.Args().Len() == 0 {
		d.showProfile()
		return nil
	} else {
		return d.switchProfile(c.Args().First())
	}

}

type ListProfileCommand struct {
}

func (d ListProfileCommand) Run(c *cli.Context) error {

	profileList := config.Profiles().List()
	term.Textf("profiles:\n")
	for _, p := range profileList {
		active := ""
		if p == config.Profiles().LoadCurrent().Name {
			active = " (active)"
		}
		term.Textf("\t- %s%s\n", p, active)

	}

	return nil
}

func (d SwitchProfileCommand) showProfile() {
	profileName := config.Profiles().LoadCurrent().Name
	term.Textf("profile: %s\n", profileName)
}

func (d SwitchProfileCommand) switchProfile(exp string) error {

	// fuzzy match
	pList := config.Profiles().List()
	sel := make([]string, 0)
	for _, p := range pList {
		if strings.HasPrefix(p, exp) {
			sel = append(sel, p)
		}
	}
	if len(sel) == 0 {
		term.Failf("no profile found for [%s]\n", exp)
		return nil
	}
	if len(sel) > 1 {
		term.Failf("multiple profiles found for [%s]: %s\n", exp, strings.Join(sel, ", "))
		return nil
	}
	profileName := sel[0]

	// switch
	err := config.Profiles().Switch(profileName)
	if err != nil {
		term.Failf("failed to switch profile: [%s]\n", err)
		return err
	}
	term.Successf("switched to profile [%s]\n", profileName)
	return nil
}
