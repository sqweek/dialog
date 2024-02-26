package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/sqweek/dialog"

	"gioui.org/app"
	"gioui.org/op"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

func main() {
	go func() {
		w := app.NewWindow(
			app.Size(300, 300),
		)
		if err := loop(w); err != nil {
			log.Fatal(err)
		}
		os.Exit(0)
	}()
	app.Main()
}

type selectedResults struct {
	File  string
	Dir   string
	Error error
}

var selected *selectedResults
var selectedMutex sync.RWMutex

func loop(w *app.Window) error {
	var ops op.Ops

	var startProcess widget.Clickable
	th := material.NewTheme()
	for {
		switch e := w.NextEvent().(type) {
		case app.DestroyEvent:
			return e.Err
		case app.FrameEvent:
			gtx := app.NewContext(&ops, e)
			if startProcess.Clicked(gtx) {
				// this needs to be run in a separate goroutine
				go func() {
					file, dir, err := runDialogs()
					selectedMutex.Lock()
					selected = &selectedResults{
						File:  file,
						Dir:   dir,
						Error: err,
					}
					selectedMutex.Unlock()
				}()
			}
			text := "Start process"
			if selected != nil {
				text = fmt.Sprintf("%s %s %+v", selected.File, selected.Dir, selected.Error)
			}

			material.Button(th, &startProcess, text).Layout(gtx)
			e.Frame(gtx.Ops)
		}
	}
}

func runDialogs() (string, string, error) {
	dialog.Message("%s", "Please select a file").Title("Hello world!").Info()
	file, err := dialog.File().Title("Save As").Filter("All Files", "*").Save()
	if err != nil {
		return "", "", err
	}
	dialog.Message("You chose file: %s", file).Title("Goodbye world!").Error()
	dir, err := dialog.Directory().Title("Now find a dir").Browse()
	if err != nil {
		return "", "", err
	}
	return file, dir, err
}
