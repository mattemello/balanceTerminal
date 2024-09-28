package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func PageError() *tview.Modal /*controll if there is sometingh better*/ {

	modal := tview.NewModal()

	modal.SetText("Not valid")
	modal.AddButtons([]string{"go back"})

	modal.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		pages.SwitchToPage("Main")
	})

	modal.SetBackgroundColor(tcell.ColorBlack)
	modal.SetTextColor(tcell.ColorGhostWhite)
	modal.SetBorderColor(tcell.ColorBlack)

	return modal
}
