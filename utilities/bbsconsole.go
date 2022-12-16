package main

import (
	ui "github.com/rickcollette/clui"
)

func main() {
	ui.InitLibrary()
	defer ui.DeinitLibrary()

	// Setup the screen layout
	view := ui.AddWindow(0, 0, 80, 24, "TVUBBS Console")
	topNav := ui.CreateFrame(view, 80, 3, ui.BorderNone, ui.Fixed)

	// Create the top navigation bar
	topStatus := ui.CreateLabel(topNav, 12, 1, "Initializing...", ui.Fixed)
	topStatus.SetSize(20, 1)
	topStatus.SetPos(1, 1)

	// Wait For Call Button
	waitForcall := ui.CreateT(topNav, 5, 1, " WFC ", ui.Fixed)
	waitForcall.SetSize(5, 1)
	waitForcall.SetPos(21, 1)
	waitForcall.onClick(func(ev ui.Event) {
		go ui.Stop()
	})

	ui.MainLoop()
}
