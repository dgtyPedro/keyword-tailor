package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	// "fyne.io/fyne/v2/dialog"
	"github.com/sqweek/dialog"
	"image/color"
	"net/url"
)

type Navigator struct {
	step     int
	lastStep *int
	window   fyne.Window
}

var tailor = app.New()

var navigator = &Navigator{
	step:     1,
	lastStep: nil,
	window:   tailor.NewWindow("Keyword Tailor"),
}

func loadFileInput() *fyne.Container {
	inputButton := widget.NewButton("Select your document to start", func() {
		filename, _ := dialog.File().Title("Select your document").Load()
		fmt.Println(filename)
		renderStep(2)
	})
	title := canvas.NewText("Keyword Tailor", color.White)
	linkURL, _ := url.Parse("https://docs.google.com/document/d/1DaJguK3Wo_Z7I177fyVfvrEPU1QJZzi0Tx4IKsBDNF8/edit?usp=sharing")

	template := widget.NewHyperlink("Get our recommended cv template here",
		linkURL)
	return container.NewVBox(
		layout.NewSpacer(),
		container.NewHBox(
			layout.NewSpacer(),
			container.NewVBox(container.NewCenter(title), inputButton),
			layout.NewSpacer(),
		),
		layout.NewSpacer(),
		container.NewCenter(template),
	)
}

func loadFileInfo() *fyne.Container {
	return container.NewVBox(
		layout.NewSpacer(),
		container.NewHBox(
			layout.NewSpacer(),
			widget.NewButtonWithIcon("Go Back", theme.NavigateBackIcon(), func() {
				renderStep(1) // Voltar para a página anterior
			}),
			layout.NewSpacer(),
		),
		layout.NewSpacer(),
	)
}

func renderStep(step int) {
	fmt.Println(step)
	navigator.lastStep = &navigator.step
	navigator.step = step
	switch step {
	case 2:
		navigator.window.SetContent(loadFileInfo())
	default:
		navigator.window.SetContent(loadFileInput())
	}
}

func main() {
	renderStep(1)
	myWindow := navigator.window
	myWindow.Resize(fyne.NewSize(900, 900))
	myWindow.ShowAndRun()
}
