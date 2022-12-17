package main

import (
	"log"
	"net/http"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

type Config struct {
	App            fyne.App
	MainWindow     fyne.Window
	InfoLog        *log.Logger
	ErrorLog       *log.Logger
	PriceContainer *fyne.Container
	Toolbar        *widget.Toolbar
	HTTPClient     *http.Client
}

var myApp Config

func main() {
	fyneApp := app.NewWithID("ca.gocode.goldwatcher.preferneces")
	myApp.App = fyneApp

	myApp.InfoLog = log.New(os.Stdin, "Info\t", log.Ldate|log.Ltime)
	myApp.InfoLog = log.New(os.Stdin, "Error\t", log.Ldate|log.Ltime|log.Lshortfile)

	myApp.MainWindow = fyneApp.NewWindow("GoldWatcher")

	myApp.makeUI()

	myApp.MainWindow.Resize(fyne.Size{Width: 770, Height: 410})
	myApp.MainWindow.SetFixedSize(true)
	myApp.MainWindow.SetMaster()

	myApp.MainWindow.CenterOnScreen()
	myApp.MainWindow.ShowAndRun()
}
