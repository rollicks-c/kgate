package cli

import (
	"github.com/rollicks-c/kgate/internal/cli/commands/forwards"
	"github.com/rollicks-c/kgate/internal/cli/commands/groups"
	"github.com/rollicks-c/kgate/internal/cli/commands/profile"
	"github.com/urfave/cli/v2"
)

func createCommands() []*cli.Command {

	cmdList := []*cli.Command{
		createForwardCommands(),
		createGroupsCommands(),
		createProfileCommands(),
	}
	return cmdList
}

func createForwardCommands() *cli.Command {
	return &cli.Command{
		Name:    "forward",
		Aliases: []string{"f"},
		Action:  forwards.Start,
		Usage:   "forwards [-a | -g <g1>, -g <g2>,...]",
		Flags: []cli.Flag{
			forwards.FlagGroup,
			forwards.FlagAll,
		},
	}
}

func createGroupsCommands() *cli.Command {
	return &cli.Command{
		Name:    "groups",
		Aliases: []string{"g"},
		Action:  groups.List,
		Subcommands: []*cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "list",
				Action:  groups.List,
			},
		},
	}
}

func createProfileCommands() *cli.Command {
	return &cli.Command{
		Name:     "profiles",
		Aliases:  []string{"p"},
		Action:   profile.Switch,
		HideHelp: true,
		Subcommands: []*cli.Command{
			{
				Name:    "list",
				Aliases: []string{"ls"},
				Usage:   "list",
				Action:  profile.List,
			},
			{
				Name:    "create",
				Aliases: []string{"c"},
				Usage:   "create <name>",
				Action:  profile.Create,
			},
			{
				Name:     "switch",
				Aliases:  []string{"s"},
				Usage:    "switch",
				Action:   profile.Switch,
				HideHelp: true,
			},
		},
	}
}
