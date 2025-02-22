package gate

import (
	"fmt"
	"github.com/rollicks-c/kgate/internal/config"
)

func RunGroups(groups ...config.PortGroup) error {

	// sanity check
	if len(groups) == 0 {
		return fmt.Errorf("no groups found")
	}

	// start session
	newController(groups...).Run()

	return nil
}
