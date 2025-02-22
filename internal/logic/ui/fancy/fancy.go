package fancy

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/rollicks-c/kgate/internal/logic/model"
	"os"
	"syscall"
)

func (f Frontend) setupKeyHandler(controller model.Controller) {

	f.app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 's' {
			controller.TogglePause()
		}
		if event.Rune() == 'q' {
			controller.Quit()
		}
		if event.Key() == tcell.KeyCtrlC {
			controller.Quit()
			_ = syscall.Kill(os.Getpid(), syscall.SIGTERM)
		}
		return event
	})

}

func (f Frontend) getStatusColor(value model.Status) tcell.Color {
	switch value {
	case model.Running:
		return tcell.ColorGreen
	case model.Stopped:
		return tcell.ColorOrange
	case model.Restart:
		return tcell.ColorYellow
	case model.Failure:
		return tcell.ColorRed
	default:
		return tcell.ColorWhite
	}
}

func (f Frontend) getStatusText(value model.Status) string {

	switch value {
	case model.Running:
		return "🟢 Running"
	case model.Stopped:
		return "⛔ Stopped"
	case model.Restart:
		return "🔄 Restarting"
	case model.Failure:
		return "❌ Failure"
	default:
		return "⚪ Unknown"
	}
}

func (f Frontend) updateTableRow(update model.Update) {

	row := update.SortIndex + 1

	f.layout.table.SetCell(row, 0,
		tview.NewTableCell(padText(update.Group, 10)).
			SetTextColor(tcell.ColorWhite).
			SetAlign(tview.AlignLeft).
			SetSelectable(true),
	)
	f.layout.table.SetCell(row, 1,
		tview.NewTableCell(padText(update.PortForward, 45)).
			SetTextColor(tcell.ColorLightBlue).
			SetAlign(tview.AlignLeft).
			SetSelectable(true),
	)
	f.layout.table.SetCell(row, 2,
		tview.NewTableCell(padText(f.getStatusText(update.Status), 20)).
			SetTextColor(f.getStatusColor(update.Status)).
			SetAlign(tview.AlignLeft).
			SetSelectable(true),
	)
	f.layout.table.SetCell(row, 3,
		tview.NewTableCell(update.Message).
			SetTextColor(f.getStatusColor(update.Status)).
			SetTextColor(f.getStatusColor(update.Status)).
			SetAlign(tview.AlignLeft).
			SetSelectable(false),
	)
}
