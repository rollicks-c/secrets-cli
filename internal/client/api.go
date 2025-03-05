package client

import (
	"github.com/urfave/cli/v2"
)

func CreateClient() *cli.App {
	app := cli.NewApp()
	app.Name = "Secrets CLI"
	app.Commands = createCommands()
	return app
}
