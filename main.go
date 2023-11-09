package main

import (
	"github.com/getlantern/systray"
)

var (
	handler = NewHandler()
)

func main() {
	go func() {
		handler.Listen()
	}()
	systray.Run(onReady, func() {})
}

func onReady() {
	systray.SetTitle("Awesome App")
	systray.SetTooltip("Lantern")

	handler.AddAction("Trim", "Remove Spaces at start and end of string", Trim)
	handler.AddAction("Capitalize Uppercase", "Capitalize each word if value is uppercase", CapitalizeUppercase)

	systray.AddSeparator()

	mQuitOrig := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-mQuitOrig.ClickedCh
		systray.Quit()
	}()
}
