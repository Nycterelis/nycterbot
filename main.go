package main

import (
	"fmt"

	"fyne.io/fyne/layout"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {
	app := app.New()

	w := app.NewWindow("NycterBot")

	// Creating the tool bar
	// toolbar := widget.NewToolbar()

	// Creating the components of the First HBox
	message := widget.NewMultiLineEntry()
	message.SetReadOnly(true)
	message.SetPlaceHolder(`This is an example of placeholding text 
Hell yeah, Resinger is drunk af. 
I cannot remember what linefeed is. 
I am the worst developer ever. Who ever is my mother. I was orphaned as a young
boy.
Can you smell-l-l-l-l-l-l-l-l what The ROCK is cookin?`)
	chatList := widget.NewMultiLineEntry()
	chatList.SetReadOnly(true)

	// Creating the Components of the Second HBox
	textEntry := widget.NewEntry()
	enterButton := widget.NewButton("Send", (func() { app.Quit() }))

	// Constructing boxes for use.
	bigTopBox := widget.NewHBox(message, chatList)
	smallBottomBox := widget.NewHBox(textEntry, enterButton)
	boxes := make([]fyne.CanvasObject, 0)
	boxes = append(boxes, bigTopBox)
	boxes = append(boxes, smallBottomBox)
	theLayout := layout.NewHBoxLayout()
	// layoutSize := fyne.NewSize(1000, 300)
	// theLayout.Layout(boxes, layoutSize)

	theContainer := fyne.NewContainerWithLayout(
		theLayout, bigTopBox)
	theContainer.AddObject(smallBottomBox)

	w.SetContent(theContainer)

	fmt.Print(w.Content().Size())

	w.ShowAndRun()
}
