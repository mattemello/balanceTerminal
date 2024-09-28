package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func PageError(m string) *tview.Modal {

	modal := tview.NewModal()

	modal.SetText(m)
	modal.AddButtons([]string{"go back"})

	modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		pages.AddAndSwitchToPage("Main", menuCreation(), true)
	})

	modal.SetBackgroundColor(tcell.ColorBlack)
	modal.SetTextColor(tcell.ColorGhostWhite)
	modal.SetBorderColor(tcell.ColorBlack)

	return modal
}
