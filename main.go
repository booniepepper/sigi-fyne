package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Sigi")

	stack := container.NewGridWithColumns(3,
		widget.NewLabel("Now"), widget.NewLabel("Make an app"), widget.NewButtonWithIcon("", theme.ConfirmIcon(), func() {}),
		widget.NewLabel("1"), widget.NewLabel("Take a nap"), widget.NewButtonWithIcon("", theme.ConfirmIcon(), func() {}),
		widget.NewLabel("2"), widget.NewLabel("Do more stuff"), widget.NewButtonWithIcon("", theme.ConfirmIcon(), func() {}))
	w.SetContent(stack)
	w.ShowAndRun()
}
