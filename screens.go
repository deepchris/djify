package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func accountPage() fyne.CanvasObject {
	var logo canvas.Image
	var logoSize fyne.Size

	logoSize = fyne.NewSize(236.2/1.75, 70.9/1.75)
	logo.File = `C:\Users\t1008257\OneDrive - Axcelis Technologies, Inc\Documents\GitHub\djify\Spotify_Logo_RGB_Green.png`
	logo.SetMinSize(logoSize)

	spotifyLogo := container.NewHBox(widget.NewCard("Spotify", "Account Login", &logo))
	return container.NewVBox(spotifyLogo, layout.NewSpacer())
}

func playlistPage() fyne.CanvasObject {
	return container.NewVBox(widget.NewLabel("Playlist Page"))
}

func seedPage() fyne.CanvasObject {
	return container.NewVBox(widget.NewLabel("Generator Settings Page"))
}

func donatePage() fyne.CanvasObject {
	return container.NewHBox(widget.NewLabel("Donations Page <3"))
}
func settingsPage() fyne.CanvasObject {
	return container.NewHBox(widget.NewLabel("Settings Page"))
}
