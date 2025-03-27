package profile

import (
	"github.com/rollicks-c/kgate/internal/config"
	"github.com/rollicks-c/term"
	"strings"
)

func showProfile() {
	profileName := config.Profiles().LoadCurrent().Name
	term.Infof("profile: %s\n", profileName)
}

func switchProfile(exp string) error {

	// fuzzy match
	list, _ := Find(exp)
	if len(list) == 0 {
		term.Failf("no profile found matching [%s]\n", exp)
		return nil
	}
	if len(list) > 1 {
		term.Failf("multiple profiles found matching [%s]: %s\n", exp, strings.Join(list, ", "))
		return nil
	}
	profileName := list[0]

	// switch
	err := config.Profiles().Switch(profileName)
	if err != nil {
		term.Failf("failed to switch profile: [%s]\n", err)
		return err
	}
	term.Successf("switched to profile [%s]\n", profileName)
	return nil
}
