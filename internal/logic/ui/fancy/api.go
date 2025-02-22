package fancy

import (
	"github.com/rivo/tview"
	"github.com/rollicks-c/kgate/internal/logic/model"
	"github.com/rollicks-c/term"
)

type Frontend struct {
	app     *tview.Application
	layout  *appLayout
	records map[string]int // id->row
}

func New() *Frontend {

	// build app
	layout := createLayout()
	app := tview.NewApplication().
		EnableMouse(false).
		SetRoot(layout.root, true)

	return &Frontend{
		app:     app,
		layout:  layout,
		records: map[string]int{},
	}
}

func (f Frontend) Run(controller model.Controller) {

	// install event handlers
	f.setupKeyHandler(controller)

	// run
	if err := f.app.Run(); err != nil {
		panic(err)
	}

}

func (f Frontend) ShowMessage(msg string) {
	f.app.QueueUpdateDraw(func() {
		f.layout.msgBox.SetText(msg + " ")
	})
}

func (f Frontend) Stop() {
	f.app.Stop()
	term.Warnf("shutting down UI\n")
	f.app.SetScreen(nil)
}

func (f Frontend) Update(update model.Update) {
	f.app.QueueUpdateDraw(func() {
		f.updateTableRow(update)
	})
}
