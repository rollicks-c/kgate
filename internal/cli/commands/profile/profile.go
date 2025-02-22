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
	pList := config.Profiles().List()
	sel := make([]string, 0)
	for _, p := range pList {
		if strings.HasPrefix(p, exp) {
			sel = append(sel, p)
		}
	}
	if len(sel) == 0 {
		term.Failf("no profile found matching [%s]\n", exp)
		return nil
	}
	if len(sel) > 1 {
		term.Failf("multiple profiles found matching [%s]: %s\n", exp, strings.Join(sel, ", "))
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
