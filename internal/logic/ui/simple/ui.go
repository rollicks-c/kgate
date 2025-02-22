package simple

import (
	"fmt"
	"github.com/rollicks-c/kgate/internal/logic/model"
	"os"
	"os/exec"
	"runtime"
)

type Frontend struct {
	procList map[string]model.Update
}

func (f Frontend) Run(controller model.Controller) {
}

func New() *Frontend {
	return &Frontend{
		procList: make(map[string]model.Update),
	}
}

func (f Frontend) ShowMessage(msg string) {
	fmt.Println(msg)
}

func (f Frontend) Stop() {
	f.clearTerminal()
}

func (f Frontend) Update(update model.Update) {
	f.clearTerminal()
	f.procList[update.ID] = update
	for _, proc := range f.procList {
		fmt.Println(proc)
	}

}

func (f Frontend) clearTerminal() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}
