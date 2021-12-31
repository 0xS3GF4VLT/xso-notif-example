package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

	labelHello := widget.NewLabel("Hello World")

	x := XSOConnect{}
	x.Connect()

	w.SetContent(container.NewVBox(
		labelHello,
		widget.NewButton("Say Hello", func() {
			x.SendNotif("", 0)
		}),
		widget.NewRichTextFromMarkdown("_(requires XS Overlay for Steam VR)_"),
	))

	w.CenterOnScreen()
	w.Resize(fyne.NewSize(600, 600))
	w.ShowAndRun()
}
