package ui

import (
	"github.com/gdamore/tcell/v2"
	errorhand "github.com/mattemello/balanceTerminal/errorHand"
	"github.com/mattemello/balanceTerminal/sqlScript"
	"github.com/rivo/tview"
)

func insertFTags() *tview.Flex {
	f := flexCreation()

	f.AddItem(insertTags(), 0, 9, true)
	f.AddItem(footSetTags(), 0, 1, false)

	return f
}

func footSetTags() *tview.TextView {
	text1 := tview.NewTextView().SetText("(esc) back \n (" + string(tcell.RuneUArrow) + "/" + string(tcell.RuneDArrow) + ") move  ").SetTextColor(tcell.ColorSnow)
	text1.SetTextAlign(tview.AlignCenter).SetBorder(true)

	return text1
}

func insertTags() *tview.Form {
	form := tview.NewForm()

	form.SetFieldBackgroundColor(tcell.Color(tcell.ColorValues[12]))
	form.SetFieldTextColor(tcell.ColorSnow)
	form.SetLabelColor(tcell.ColorWhiteSmoke)
	form.SetButtonBackgroundColor(tcell.Color(tcell.ColorValues[12]))

	form.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {

		if event.Key() == tcell.KeyEsc {
			pages.SwitchToPage("Main")
		}

		if event.Key() == 258 {
			SwitchFocus(form, 1)
		} else if event.Key() == 257 {
			SwitchFocus(form, -1)
		}

		return event
	})

	var fiel string

	form.SetBorder(true).SetTitle("Page to add a new tags")

	form.AddInputField("Tag: ", "", 20, nil, func(text string) {
		fiel = text
	})

	form.AddButton("Save tag", func() {

		for _, c := range sqlScript.AllTags {
			if c == fiel {
				pages.AddAndSwitchToPage("err", PageError("The tag is already in"), true)
			}
		}

		err := sqlScript.SaveTags(fiel)

		if err != nil {
			errorhand.BadSaving(err)
			pages.AddAndSwitchToPage("err", PageError("Error whit the save, controll the log file"), true)
		} else {
			sqlScript.SaveTag(fiel)
			pages.RemovePage("menu")
			pages.AddAndSwitchToPage("Main", menuCreation(), true)
		}

	})

	return form
}
