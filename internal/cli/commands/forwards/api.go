package forwards

import (
	"github.com/urfave/cli/v2"
)

var (
	FlagGroup = &cli.StringSliceFlag{
		Name:    "group",
		Aliases: []string{"g"},
	}
	FlagAll = &cli.BoolFlag{
		Name:    "all",
		Value:   false,
		Aliases: []string{"a"},
	}
)

func Start(c *cli.Context) error {
	selectedGroups := FlagGroup.Get(c)
	allGroups := FlagAll.Get(c)
	return startGroups(allGroups, selectedGroups...)
}

func StartAll(c *cli.Context) error {
	return startGroups(true)
}
