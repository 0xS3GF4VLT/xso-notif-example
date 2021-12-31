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

	accordion := widget.NewAccordion()
	accordion.Append(widget.NewAccordionItem("Expando", container.NewVBox(
		widget.NewRichTextFromMarkdown("## Indeed, we rich"),
	)))

	x := XSOConnect{}

	x.Connect()

	w.SetContent(container.NewVBox(
		labelHello,
		widget.NewButton("Sup", func() {
			labelHello.SetText("Yuh!")
			x.SendNotif("", 0)
		}),
		container.NewAppTabs(
			container.NewTabItem("Thing 1", container.NewVBox(
				widget.NewLabel("yuh"),
			)),
			container.NewTabItem("Thing 2", container.NewVBox(
				widget.NewLabel("Life, uhhh"),
				widget.NewRichTextFromMarkdown("__... finds away__"),
			)),
		),
		accordion,
	))

	w.CenterOnScreen()
	w.Resize(fyne.NewSize(600, 600))
	w.ShowAndRun()
}
