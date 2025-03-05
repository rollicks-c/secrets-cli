package alias

import (
	"github.com/rollicks-c/secrets-cli/internal/client/commands/params"
	"github.com/rollicks-c/secrets-cli/internal/config"
	"github.com/rollicks-c/term"
	"github.com/urfave/cli/v2"
	"strings"
)

func Run(c *cli.Context) error {
	// collect params
	var pc = params.NewCollector(c.Args().Slice())
	if err := pc.Validate(1); err != nil {
		return err
	}
	searchTerms := pc.GetBefore("-")
	cmdArgs := pc.GetAfter("-")
	name := ""
	if len(searchTerms) == 1 {
		name = searchTerms[0]
		searchTerms = searchTerms[1:]
	}

	// resolve alias
	alias, err := resolveAlias(name, searchTerms...)
	if err != nil {
		return err
	}

	// invoke
	cmd := "null " + alias.Command
	args := strings.Split(cmd, " ")
	args = append(args, cmdArgs...)
	if err := c.App.Run(args); err != nil {
		return err
	}

	return nil

}

func Add(c *cli.Context) error {

	// collect params
	var pc = params.NewCollector(c.Args().Slice())
	if err := pc.Validate(2); err != nil {
		return err
	}
	name := pc.GetFirst()
	cmd := pc.GetBefore("-")
	cmd = cmd[1:]
	tags := pc.GetAfter("-")

	// add alias
	conf := config.Profiles().LoadCurrent()
	aliasMap := conf.Data.Aliases
	if _, ok := aliasMap[name]; ok {
		term.Warnf("alias name %s already exists\n", name)
		return nil
	}
	aliasMap[name] = config.Alias{
		Command: strings.Join(cmd, " "),
		Tags:    strings.Join(tags, " "),
	}
	conf.Data.Aliases = aliasMap
	config.Profiles().Update(conf)

	// report
	term.Successf("alias %s added\n", name)
	return nil
}

func Remove(c *cli.Context) error {

	// collect params
	var pc = params.NewCollector(c.Args().Slice())
	if err := pc.Validate(1); err != nil {
		return err
	}
	alias := pc.GetString(0)

	// add alias
	conf := config.Profiles().LoadCurrent()
	aliasMap := conf.Data.Aliases
	if _, ok := aliasMap[alias]; !ok {
		term.Failf("alias %s does not exists\n", alias)
		return nil
	}
	delete(aliasMap, alias)
	conf.Data.Aliases = aliasMap
	config.Profiles().Update(conf)

	// report
	term.Warnf("alias %s removed\n", alias)
	return nil
}

func List(c *cli.Context) error {

	aliasList := config.Profiles().LoadCurrent().Data.Aliases
	term.Infof("aliases:\n")
	for a, p := range aliasList {
		term.Infof("\t%s\t= %s\n", a, p)
	}

	return nil
}
