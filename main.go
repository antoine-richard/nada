package main

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	width, height = 640, 480
)

var (
	x, y float64 = width / 2, height / 2
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Nada",
		Bounds: pixel.R(0, 0, width, height),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd := imdraw.New(nil)

	imd.Color = colornames.Black
	imd.Push(pixel.V(200, 0), pixel.V(200, height))
	imd.Line(1)

	imd.Color = colornames.Red
	imd.Push(pixel.V(x, y))
	imd.Circle(5, 0)

	for !win.Closed() {
		win.Clear(colornames.Skyblue)
		imd.Draw(win)
		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
