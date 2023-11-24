package main

import (
	"github.com/getlantern/systray"
)

var Version string
var handler = NewHandler()

func main() {
	go func() {
		handler.Listen()
	}()
	systray.Run(onReady, func() {})
}

func onReady() {
	data, err := Asset("winres/taskbaricon.ico")
	if err != nil {
		panic("Taskbar icon data not found in binary")
	}

	systray.SetTemplateIcon(data, data)

	systray.SetTitle("Cliptr")
	systray.SetTooltip("Cliptr " + Version)

	handler.AddAction("Trim", "Remove Spaces at start and end of string", Trim)
	handler.AddAction("Capitalize Uppercase", "Capitalize each word if value is uppercase", CapitalizeUppercase)

	systray.AddSeparator()

	mQuitOrig := systray.AddMenuItem("Quit", "Close the application")
	go func() {
		<-mQuitOrig.ClickedCh
		systray.Quit()
	}()
}
