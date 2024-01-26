package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"os/exec"
	"runtime"

	"gioui.org/app"
	"gioui.org/io/system"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"github.com/sqweek/dialog"
)

type Navigator struct {
	step int
}

var nav = Navigator{
	step: 1,
}

func gio() {
	go func() {
		w := app.NewWindow()
		err := run(w)
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

func run(w *app.Window) error {
	th := material.NewTheme()
	var ops op.Ops

	var startButton widget.Clickable
	var docRedirect widget.Clickable

	for {
		switch e := w.NextEvent().(type) {
		case system.DestroyEvent:
			return e.Err
		case system.FrameEvent:
			gtx := layout.NewContext(&ops, e)

			if startButton.Clicked(gtx) {
				filepath, err := dialog.File().Title("Select your document").Load()
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println(filepath)
			}

			if docRedirect.Clicked(gtx) {
				fmt.Println("redirecting")
				var err error
				url := "https://docs.google.com/document/d/1DaJguK3Wo_Z7I177fyVfvrEPU1QJZzi0Tx4IKsBDNF8"
				switch runtime.GOOS {
				case "linux":
					err = exec.Command("xdg-open", url).Start()
				case "windows":
					err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
				case "darwin":
					err = exec.Command("open", url).Start()
				default:
					err = fmt.Errorf("unsupported platform")
				}
				if err != nil {
					log.Fatal(err)
				}
			}

			layout.Flex{
				Axis:      layout.Vertical,
				Spacing:   layout.SpaceSides,
				Alignment: layout.Middle,
			}.Layout(gtx,
				layout.Flexed(2,
					func(gtx layout.Context) layout.Dimensions {
						return layout.Flex{
							Axis:      layout.Vertical,
							Spacing:   layout.SpaceSides,
							Alignment: layout.Middle,
						}.Layout(gtx,
							layout.Rigid(
								func(gtx layout.Context) layout.Dimensions {
									t := material.Body1(th, "Keyword Taylor")
									t.Alignment = text.Middle
									return t.Layout(gtx)
								},
							),
							layout.Rigid(
								layout.Spacer{Height: unit.Dp(25)}.Layout,
							),
							layout.Rigid(
								func(gtx layout.Context) layout.Dimensions {
									gtx.Constraints = layout.Exact(image.Pt(300, 50))
									btn := material.Button(th, &startButton, "Insert your document to start")
									return btn.Layout(gtx)
								},
							),
						)
					},
				),
				layout.Flexed(0.1,
					func(gtx layout.Context) layout.Dimensions {
						return layout.Stack{}.Layout(gtx,
							layout.Stacked(
								func(gtx layout.Context) layout.Dimensions {
									return material.Clickable(gtx, &docRedirect, func(gtx layout.Context) layout.Dimensions {
										return material.Subtitle1(th, "Need a cv? Use our template.").Layout(gtx)
									})
								},
							),
						)
					},
				),
			)

			e.Frame(gtx.Ops)
		}
	}
}
