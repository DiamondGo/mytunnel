package main

import (
	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/container"
)

func getMainWindow(parent fyne.Window) *fyne.Container {

	configTopAera := getConfigWindow(parent)
	content := container.NewBorder(configTopAera, nil, nil, nil, getTunnelWindow(parent))
	return content
}