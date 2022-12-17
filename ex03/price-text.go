package main

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func (app *Config) getPriceText() (*canvas.Text, *canvas.Text, *canvas.Text) {
	var g Gold

	var open, current, change *canvas.Text

	gold, err := g.GetPrices()

	if err != nil {
		grey := color.NRGBA{R: 155, G: 155, B: 155, A: 255}
		open = canvas.NewText("open: Unreachable", grey)
		current = canvas.NewText("currnet: Unreachable", grey)
		change = canvas.NewText("change: Unreachable", grey)
	} else {
		displayColor := color.NRGBA{R: 0, G: 180, B: 0, A: 255}

		if gold.Price < gold.PreviousClose {
			displayColor = color.NRGBA{R: 180, G: 0, B: 0, A: 255}

		}

		openText := fmt.Sprintf("open: $%.4f %s", gold.PreviousClose, currenCy)
		currentText := fmt.Sprintf("current: $%.4f %s", gold.Price, currenCy)
		changeText := fmt.Sprintf("change: $%.4f %s", gold.Change, currenCy)

		open = canvas.NewText(openText, displayColor)
		current = canvas.NewText(currentText, displayColor)
		change = canvas.NewText(changeText, displayColor)
	}

	open.Alignment = fyne.TextAlignLeading
	current.Alignment = fyne.TextAlignLeading
	change.Alignment = fyne.TextAlignLeading

	return open, current, change
}
