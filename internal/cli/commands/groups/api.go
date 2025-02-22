package groups

import (
	"github.com/rollicks-c/kgate/internal/config"
	"github.com/rollicks-c/term"
	"github.com/urfave/cli/v2"
)

func List(c *cli.Context) error {

	profile := config.Profiles().LoadCurrent()
	term.Infof("profile [%s]\ngroups:\n", profile.Name)
	for _, g := range profile.Data.Groups {
		term.Infof("    - %s [%s]\n", g.Name, g.Target.K8sContext)
		for _, pf := range g.PortForwards {
			term.Infof("        - %s:%s/%s:%s\n", pf.LocalPort, pf.Namespace, pf.Service, pf.RemotePort)
		}
	}

	return nil
}
