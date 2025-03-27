package profile

import (
	"fmt"
	"github.com/rollicks-c/kgate/internal/config"
	"github.com/rollicks-c/term"
	"github.com/urfave/cli/v2"
	"strings"
)

func Switch(c *cli.Context) error {

	// no args: show current profile
	if c.Args().Len() == 0 {
		showProfile()
		return nil
	}

	// switch
	if err := switchProfile(c.Args().First()); err != nil {
		return err
	}

	return nil
}

func List(c *cli.Context) error {

	profileList := config.Profiles().List()
	term.Infof("profiles:\n")
	for _, p := range profileList {
		active := ""
		if p == config.Profiles().LoadCurrent().Name {
			active = " (active)"
		}
		term.Infof("\t- %s%s\n", p, active)

	}

	return nil
}

func Create(c *cli.Context) error {

	// gather profile name
	profileName := c.Args().First()
	if profileName == "" {
		return fmt.Errorf("profile name is required")
	}

	// avoid name collision
	profileList := config.Profiles().List()
	term.Infof("profiles:\n")
	for _, p := range profileList {

		if p == profileName {
			return fmt.Errorf("profile %s already exists\n", profileName)
		}
	}

	// create profile
	template := config.Profiles().LoadCurrent()
	template.Name = profileName
	template.Data.Groups = append(template.Data.Groups, config.PortGroup{
		Target: config.Target{
			K8sConfigFile: "${HOME}/.kube/config",
			K8sContext:    "context1",
		},
		PortForwards: []config.PortForward{
			{
				Namespace:  "namespace1",
				Service:    "service1",
				LocalPort:  "8080",
				RemotePort: "8080",
			},
		},
		Name: "group1",
	})
	config.Profiles().Update(template)

	term.Infof("profile %s created\n", profileName)

	return nil
}

func Find(exp string) ([]string, bool) {

	// fuzzy match
	pList := config.Profiles().List()
	sel := make([]string, 0)
	for _, p := range pList {
		if strings.HasPrefix(p, exp) {
			sel = append(sel, p)
		}
	}
	if len(sel) == 0 {
		return sel, false
	}
	if len(sel) > 1 {
		return sel, false
	}
	return sel, true
}
