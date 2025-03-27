package forwards

import (
	"fmt"
	"github.com/rollicks-c/kgate/internal/cli/commands/profile"
	"github.com/rollicks-c/kgate/internal/config"
	"github.com/urfave/cli/v2"
)

var (
	FlagProfile = &cli.StringFlag{
		Name:    "profile",
		Aliases: []string{"p"},
	}
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
	prof := config.Profiles().LoadCurrent()
	return startGroups(prof, allGroups, selectedGroups...)
}

func StartAll(c *cli.Context) error {
	profExp := FlagProfile.Get(c)
	prof := config.Profiles().LoadCurrent()
	if profExp != "" {
		list, ok := profile.Find(profExp)
		if !ok {
			return fmt.Errorf("no profile found uniquely matching [%s]", profExp)
		}
		prof = config.Profiles().Load(list[0], false)
	}
	return startGroups(prof, true)
}
