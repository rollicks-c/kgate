package cli

import (
	"github.com/rollicks-c/kgate/internal/cli/commands/forwards"
	"github.com/rollicks-c/kgate/internal/config"
	"github.com/urfave/cli/v2"
)

func CreateClient() *cli.App {
	app := cli.NewApp()
	app.Name = config.AppName
	app.Usage = config.Usage
	app.Commands = createCommands()
	app.Action = forwards.StartAll
	return app
}
