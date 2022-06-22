package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
)

const appScreenWidth = 800
const appScreenHeight = 600

var guiApp fyne.App

func main() {
	guiApp = app.New()
	myWindow := guiApp.NewWindow("DJ-ify: A DJ-friendly spotify playlist creator.")
	myWindow.Resize(fyne.NewSize(appScreenWidth, appScreenHeight))
	// myWindow.SetIcon(resourceLogoPng)

	fmt.Printf("Starting main loop\n")

	tabs := container.NewAppTabs(
		container.NewTabItem("Spotify Account", accountPage()),
		container.NewTabItem("Playlist", playlistPage()),
		container.NewTabItem("Generator Settings", seedPage()),
		container.NewTabItem("Donate", donatePage()),
		container.NewTabItem("Settings", settingsPage()),
	)

	tabs.SetTabLocation(container.TabLocationTop)

	myWindow.SetContent(tabs)

	myWindow.ShowAndRun()
	fmt.Println("Closing")

}
