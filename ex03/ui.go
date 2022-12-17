package main

import (
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

func (app *Config) makeUI() {
	openPrice, currentPrice, priceChange := app.getPriceText()

	priceContent := container.NewGridWithColumns(3, openPrice, currentPrice, priceChange)

	app.PriceContainer = priceContent

	toolBar := app.getToolBar()
	app.Toolbar = toolBar

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Prices", theme.HomeIcon(), canvas.NewText("Price", nil)),
		container.NewTabItemWithIcon("Holding", theme.InfoIcon(), canvas.NewText("Holding", nil)),
	)

	tabs.SetTabLocation(container.TabLocationTop)

	finalContent := container.NewVBox(priceContent, toolBar, tabs)

	app.MainWindow.SetContent(finalContent)

}
