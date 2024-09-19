package ui

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

var pages = tview.NewPages()

func AppCreation() *tview.Application {

	app := tview.NewApplication()
	return app
}

func PageCreation() *tview.Pages {

	textA := tview.NewTextView().SetText("(q) for quit").SetTextColor(tcell.ColorSnow)
	textA.SetBackgroundColor(tcell.ColorBlack)

	pages.AddPage("Main", menuCreation(), true, true)
	pages.AddPage("Insert", insertCreation(), true, false)
	pages.AddPage("Add", addMoney(), true, false)

	return pages
}

func SwitchFocus(form *tview.Form, val int) {
	id, bid := form.GetFocusedItemIndex()

	if bid != -1 {
		form.SetFocus(form.GetFormItemCount() - 2)
	}

	form.SetFocus(id + val)
}

func flexCreation() *tview.Flex {

	flex := tview.NewFlex().SetDirection(tview.FlexRow)

	return flex

}
