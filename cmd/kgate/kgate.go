package main

import (
	"github.com/rollicks-c/kgate/internal/cli"
	"github.com/rollicks-c/term"
	"os"
)

func main() {

	if err := cli.CreateClient().Run(os.Args); err != nil {
		term.Failf("%s\n", err.Error())
	}

}
