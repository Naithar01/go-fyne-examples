package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type App struct {
	output *widget.Label
}

var myApp App

func main() {
	a := app.New()
	w := a.NewWindow("Hello World")

	output, entry, button := myApp.makeUI()

	w.SetContent(container.NewVBox(output, entry, button))

	w.Resize(fyne.Size{Width: 500, Height: 500})
	w.ShowAndRun()
}

func (a App) makeUI() (*widget.Label, *widget.Entry, *widget.Button) {
	output := widget.NewLabel("Hello world (Label)")
	entry := widget.NewEntry()
	btn := widget.NewButton("Change Label", func() {
		if len(entry.Text) > 0 {
			a.output.SetText(entry.Text)
		}
	})

	a.output = output

	return output, entry, btn
}
