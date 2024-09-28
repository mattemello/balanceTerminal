package ui

import (
	"github.com/gdamore/tcell/v2"
	//errorhand "github.com/mattemello/balanceTerminal/errorHand"
	"github.com/rivo/tview"
)

var pages = tview.NewPages()
var app = tview.NewApplication()

func stopApp() {
	app.Stop()
}

func AppCreation() *tview.Application {

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		return event
	})
	return app
}

func PageCreation() *tview.Pages {

	textA := tview.NewTextView().SetText("(q) for quit").SetTextColor(tcell.ColorSnow)
	textA.SetBackgroundColor(tcell.ColorBlack)

	pages.AddPage("Main", menuCreation(), true, true)

	return pages
}

func SwitchFocus(form *tview.Form, val int) {
	id, bid := form.GetFocusedItemIndex()

	if bid != -1 {
		//errorhand.Controll(form.GetFormItemIndex("Insert money to add: "))
		form.SetFocus(form.GetFormItemCount() - 2)
	}

	form.SetFocus(id + val)
}

func flexCreation() *tview.Flex {

	flex := tview.NewFlex().SetDirection(tview.FlexRow)

	return flex

}
