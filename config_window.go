package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"

	//"fyne.io/fyne/v2/dialog"
	//"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func getConfigWindow(parent fyne.Window) *fyne.Container {
	path := binding.NewString()
	fileDialog := dialog.NewFileOpen(func(uri fyne.URIReadCloser, e error){
		if uri != nil {
			path.Set(uri.URI().Path())
		}
	}, parent)
	pathLabel := widget.NewEntryWithData(path)
	pathLabel.SetPlaceHolder("Select ssh executable...")
	selectButton := widget.NewButton("Browse", func(){
		/*
		dialog.NewFileOpen(func(uri fyne.URIReadCloser, e error){
		}, nil)
		*/
		fileDialog.Show()
		pathLabel.SetText("Clicked!")
	})

	//content := container.New(layout.NewHBoxLayout(), fileLable., fileSelect)
	content := container.NewBorder(nil, nil, nil, selectButton, pathLabel)

	return content
}