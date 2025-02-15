package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	app := tview.NewApplication()

	footer := struct {
		view        *tview.TextView
		fixedSize   int
		proportions int
		focus       bool
	}{
		view:        tview.NewTextView().SetTextAlign(tview.AlignLeft).SetText("Press esc to quit"),
		fixedSize:   1,
		proportions: 1,
		focus:       false,
	}

	layout := tview.NewFlex().
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Breakpoints"), 0, 1, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Debugging"), 0, 3, false)

	baseLayout := tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(layout, 0, 1, false).
		AddItem(footer.view, footer.fixedSize, footer.proportions, footer.focus)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEsc {
			app.Stop()
		}
		return event
	})
	if err := app.SetRoot(baseLayout, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
