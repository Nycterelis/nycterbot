package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	app := app.New()

	w := app.NewWindow("NycterBot")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("Dummy Label 1"),
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	))
	w.ShowAndRun()
}
