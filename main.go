package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"image/color"
	"os"
	"time"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Time")

	text := canvas.NewText(time.Now().Format("15:04:05"), color.Gray16{Y: 0x7fff})
	text.TextSize = 100
	text.Alignment = fyne.TextAlignCenter
	content := container.New(layout.NewStackLayout(), text)
	go func() {
		for range time.Tick(time.Second) {
			text.Text = time.Now().Format("15:04:05")
			text.Refresh()
		}
	}()
	if desk, ok := myApp.(desktop.App); ok {
		m := fyne.NewMenu("Clock",
			fyne.NewMenuItem("Change", func() {
				if text.Color == color.Black {
					text.Color = color.White
				} else if text.Color == color.White {
					text.Color = color.Gray16{Y: 0x7fff}
				} else {
					text.Color = color.Black
				}
				text.Refresh()
			}), fyne.NewMenuItem("Small", func() {
				text.TextSize = text.TextSize / 1.5
				text.Refresh()
			}), fyne.NewMenuItem("Big", func() {
				text.TextSize = text.TextSize * 1.5
				text.Refresh()
			}))
		desk.SetSystemTrayMenu(m)
	}
	myWindow.SetContent(content)
	myWindow.SetCloseIntercept(func() {
		os.Exit(0)
	})
	myWindow.ShowAndRun()

}
