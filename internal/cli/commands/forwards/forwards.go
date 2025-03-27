package forwards

import (
	"fmt"
	"github.com/rollicks-c/configcove/profiles"
	"github.com/rollicks-c/kgate/internal/config"
	"github.com/rollicks-c/kgate/internal/logic/gate"
)

func startGroups(profile profiles.Profile[config.Config], allGroups bool, selectedGroups ...string) error {

	// sanity check
	if allGroups && len(selectedGroups) > 0 {
		return fmt.Errorf("cannot use -%s and -%s together", FlagAll.Name, FlagGroup.Name)
	}

	// load data
	var groups []config.PortGroup
	if allGroups {
		groups = profile.Data.Groups
	} else {
		groups = filterGroups(profile.Data.Groups, selectedGroups)
	}

	// start forwards
	if err := gate.RunGroups(groups...); err != nil {
		return err
	}
	return nil
}

func filterGroups(pool []config.PortGroup, selectedNames []string) []config.PortGroup {
	var res []config.PortGroup
	for _, g := range pool {
		if contains(selectedNames, g.Name) {
			res = append(res, g)
		}
	}
	return res
}

func contains(groups []string, name string) bool {
	for _, g := range groups {
		if g == name {
			return true
		}
	}
	return false

}
