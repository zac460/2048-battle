package main

import (
	game "github.com/z-riley/go-2048-battle"
	"github.com/z-riley/go-2048-battle/debug"
	"github.com/z-riley/go-2048-battle/screen"
	"github.com/z-riley/turdgl"
)

func main() {
	// Create window
	win, err := turdgl.NewWindow(turdgl.WindowCfg{
		Title:  "2048 Battle",
		Width:  game.Width,
		Height: game.Height,
	})
	if err != nil {
		panic(err)
	}
	defer win.Destroy()

	// Register window-level keybinds (for development only)
	win.RegisterKeybind(turdgl.KeyEscape, func() { win.Quit() })
	win.RegisterKeybind(turdgl.KeyLCtrl, func() { win.Quit() })

	// Create screens
	screens := map[screen.Screen]screen.Updater{
		screen.Title:           screen.NewTitleScreen(win),
		screen.Singleplayer:    screen.NewSingleplayerScreen(win),
		screen.MultiplayerMenu: screen.NewMultiplayerMenuScreen(win),
		screen.MultiplayerJoin: screen.NewMultiplayerJoinScreen(win),
	}
	debugWidget := debug.NewDebugWidget(win)

	// Main game loop
	for win.IsRunning() {
		// Update screen
		screens[screen.CurrentScreen()].Update()

		if game.Debug {
			// Add debug overlay
			debugWidget.Update()
		}

		win.Update()
	}
}
