package main

import (
	ui "github.com/rickcollette/clui"
)

func createView() {
	// Setup the screen layout
	view := ui.AddWindow(0, 0, 80, 24, "TVUBBS Console")
	topNav := ui.CreateFrame(view, 80, 3, ui.BorderNone, ui.Fixed)

	// Create the top navigation bar
	topStatus := ui.CreateLabel(topNav, 12, 1, "Initializing...", ui.Fixed)
	topStatus.SetSize(20, 1)
	topStatus.SetPos(1, 1)

	// Wait For Call Button
	wfcButton := ui.CreateTextButton(topNav, 5, 1, " WFC ", ui.Fixed)
	wfcButton.SetSize(5, 1)
	wfcButton.SetPos(21, 1)
	wfcButton.OnClick(func(ev ui.Event) {
		go ui.Stop()
	})
	ui.SetCurrentTheme("tvubbs")

	wfcButton.Enabled()
	ui.ActivateControl(view, wfcButton)
	ui.PutEvent(ui.Event{Type: ui.EventRedraw})
}

func mainLoop() {
	// Every application must create a single Composer and
	// call its intialize method
	ui.InitLibrary()
	defer ui.DeinitLibrary()

	createView()

	// start event processing loop - the main core of the library
	ui.MainLoop()
}

func main() {
	mainLoop()
}
