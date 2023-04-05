package main

import (
	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func getTunnelWindow(parent fyne.Window) *fyne.Container {
	type tunnelInfo struct {
		enabled bool
		command string 
	}

	tunnels := []tunnelInfo{
		{enabled: true, command: "ssh pi4"},
	}

	tunnelList := widget.NewList(
		func() int {
			return len(tunnels)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(i widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(tunnels[i].command)
		},
	)

	content := container.NewBorder(tunnelList, nil, nil, nil)
	return content
}