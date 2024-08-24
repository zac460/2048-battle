package screen

import (
	"fmt"

	game "github.com/z-riley/go-2048-battle"
	"github.com/z-riley/go-2048-battle/backend"
	"github.com/z-riley/go-2048-battle/backend/grid"
	"github.com/z-riley/go-2048-battle/common"
	"github.com/z-riley/turdgl"
)

type SingleplayerScreen struct {
	win     *turdgl.Window
	backend *backend.Game

	debugGridText  *turdgl.Text
	debugTimeText  *turdgl.Text
	debugScoreText *turdgl.Text
}

// NewSingleplayerScreen constructs a new singleplayer menu screen.
func NewSingleplayerScreen(win *turdgl.Window) *SingleplayerScreen {
	return &SingleplayerScreen{
		win:     win,
		backend: backend.NewGame(),

		debugGridText:  turdgl.NewText("grid", turdgl.Vec{X: 100, Y: 100}, game.FontPath),
		debugTimeText:  turdgl.NewText("time", turdgl.Vec{X: 500, Y: 100}, game.FontPath),
		debugScoreText: turdgl.NewText("score", turdgl.Vec{X: 400, Y: 100}, game.FontPath),
	}
}

// Init initialises the screen.
func (s *SingleplayerScreen) Init() {
	s.debugGridText.SetText(s.backend.Arena.Debug())
	s.debugTimeText.SetText(s.backend.Timer.Time.String())
	s.debugScoreText.SetText(fmt.Sprint(s.backend.Score.Current))

	s.win.RegisterKeybind(turdgl.KeyUp, turdgl.KeyPress, func() {
		s.backend.ExecuteMove(grid.DirUp)
		s.debugGridText.SetText(s.backend.Arena.Debug())
	})
	s.win.RegisterKeybind(turdgl.KeyDown, turdgl.KeyPress, func() {
		s.backend.ExecuteMove(grid.DirDown)
		s.debugGridText.SetText(s.backend.Arena.Debug())
	})
	s.win.RegisterKeybind(turdgl.KeyLeft, turdgl.KeyPress, func() {
		s.backend.ExecuteMove(grid.DirLeft)
		s.debugGridText.SetText(s.backend.Arena.Debug())
	})
	s.win.RegisterKeybind(turdgl.KeyRight, turdgl.KeyPress, func() {
		s.backend.ExecuteMove(grid.DirRight)
		s.debugGridText.SetText(s.backend.Arena.Debug())
	})
	s.win.RegisterKeybind(turdgl.KeyR, turdgl.KeyPress, func() {
		s.backend.Reset()
	})
}

// Deinit deinitialises the screen.
func (s *SingleplayerScreen) Deinit() {
	s.backend.Timer.Pause()

	if err := s.backend.Save(); err != nil {
		panic(err)
	}

	s.win.UnregisterKeybind(turdgl.KeyUp, turdgl.KeyPress)
	s.win.UnregisterKeybind(turdgl.KeyDown, turdgl.KeyPress)
	s.win.UnregisterKeybind(turdgl.KeyLeft, turdgl.KeyPress)
	s.win.UnregisterKeybind(turdgl.KeyRight, turdgl.KeyPress)

}

// Update updates and draws the singleplayer screen.
func (s *SingleplayerScreen) Update() {
	s.win.SetBackground(common.BackgroundColour)

	s.debugGridText.SetText(s.backend.Arena.Debug())
	s.debugTimeText.SetText(s.backend.Timer.Time.String())
	s.debugScoreText.SetText(
		fmt.Sprint(s.backend.Score.Current, "|", s.backend.Score.High),
	)

	b, err := s.backend.Serialise()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

	s.win.Draw(s.debugGridText)
	s.win.Draw(s.debugTimeText)
	s.win.Draw(s.debugScoreText)
}
