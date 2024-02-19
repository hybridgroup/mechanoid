package main

import (
	"strconv"

	"github.com/aykevl/board"
	"github.com/aykevl/tinygl"
	"github.com/aykevl/tinygl/style"
	"github.com/aykevl/tinygl/style/basic"
	"tinygo.org/x/drivers/pixel"
)

type DisplayDevice[T pixel.Color] struct {
	Display  board.Displayer[T]
	Screen   *tinygl.Screen[T]
	Theme    *basic.Basic[T]
	VBox     *tinygl.VBox[T]
	Header   *tinygl.Text[T]
	PingText *tinygl.Text[T]
	PongText *tinygl.Text[T]
}

func NewDisplayDevice[T pixel.Color](disp board.Displayer[T]) DisplayDevice[T] {
	// Determine size and scale of the screen.
	width, height := disp.Size()
	scalePercent := board.Display.PPI() * 100 / 120

	// Initialize the screen.
	buf := pixel.NewImage[T](int(width), int(height)/4)
	screen := tinygl.NewScreen[T](disp, buf, board.Display.PPI())
	theme := basic.NewTheme(style.NewScale(scalePercent), screen)
	header := theme.NewText("Hello, TinyWASM")
	pingText := theme.NewText("Ping")
	pongText := theme.NewText("Pong")

	vbox := theme.NewVBox(header, pingText, pongText)
	screen.SetChild(vbox)
	screen.Update()
	board.Display.SetBrightness(board.Display.MaxBrightness())

	return DisplayDevice[T]{
		Display:  disp,
		Screen:   screen,
		Theme:    theme,
		VBox:     vbox,
		Header:   header,
		PingText: pingText,
		PongText: pongText,
	}
}

func (d *DisplayDevice[T]) Ping(count int) {
	d.PingText.SetText("Ping " + strconv.Itoa(count))
	d.Screen.Update()
}

func (d *DisplayDevice[T]) Pong(count int) {
	d.PongText.SetText("Pong " + strconv.Itoa(count))
	d.Screen.Update()
}
