package fancy

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/rollicks-c/kgate/internal/config"
)

type appLayout struct {
	root   *tview.Flex
	msgBox *tview.TextView
	table  *tview.Table
}

func createLayout() *appLayout {

	// create controls
	header, msgBox := createHeader()
	table := createTable()

	// setup appLayout
	flex := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(header, 3, 1, false).
		AddItem(table, 0, 1, true)

	return &appLayout{
		root:   flex,
		msgBox: msgBox,
		table:  table,
	}

}

func createHeader() (*tview.Flex, *tview.TextView) {

	// create controls
	titleView := tview.NewTextView().
		SetText(fmt.Sprintf("[yellow::b] %s", config.AppName)).
		SetDynamicColors(true).
		SetTextAlign(tview.AlignLeft)

	msgView := tview.NewTextView().
		SetText("").
		SetDynamicColors(true).
		SetTextAlign(tview.AlignRight)

	// appLayout
	headerFlex := tview.NewFlex().
		SetDirection(tview.FlexColumn).
		AddItem(titleView, 0, 1, false).
		AddItem(nil, 0, 2, false).
		AddItem(msgView, 50, 0, false)
	headerFlex.SetBorder(true)

	return headerFlex, msgView

}

func createTable() *tview.Table {

	// create table
	table := tview.NewTable().
		SetBorders(false).
		SetSelectable(false, false)

	// add headers
	headers := []string{"Group", "Port Forward", "Status", "Info"}
	widths := []int{15, 15, 10, 25} // Column width for padding
	for i, header := range headers {
		table.SetCell(0, i,
			tview.NewTableCell(padText(header, widths[i])).
				SetTextColor(tcell.ColorBlack).
				SetBackgroundColor(tcell.ColorWhite).
				SetAlign(tview.AlignLeft).
				SetSelectable(false).
				SetStyle(tcell.StyleDefault.Bold(true)),
		)
	}

	return table
}

func padText(text string, width int) string {
	if len(text) >= width {
		return text[:width]
	}
	return fmt.Sprintf("%-*s", width, text)
}
