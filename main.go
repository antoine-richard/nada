package main

import (
	"math"
	"time"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

const (
	windowWidth, windowHeight = 640, 480
	rotationSpeed             = 2
	walkSpeed                 = 20
)

type player struct {
	x, y float64
	dir  float64
}

var p = player{
	x:   windowWidth / 2,
	y:   windowHeight / 2,
	dir: math.Pi / 2,
}

func update(win *pixelgl.Window, dt float64) {
	if win.Pressed(pixelgl.KeyLeft) {
		p.dir += rotationSpeed * dt
	}
	if win.Pressed(pixelgl.KeyRight) {
		p.dir -= rotationSpeed * dt
	}
	if win.Pressed(pixelgl.KeyDown) {
		p.x -= walkSpeed * dt * math.Cos(p.dir)
		p.y -= walkSpeed * dt * math.Sin(p.dir)
	}
	if win.Pressed(pixelgl.KeyUp) {
		p.x += walkSpeed * dt * math.Cos(p.dir)
		p.y += walkSpeed * dt * math.Sin(p.dir)
	}
}

func drawMinimap(t pixel.Target) {
	imd := imdraw.New(nil)

	// player position
	imd.Color = colornames.Red
	imd.Push(pixel.V(p.x, p.y))
	imd.Circle(3, 0)

	// player direction
	imd.Push(pixel.V(p.x, p.y), pixel.V(p.x+math.Cos(p.dir)*15, p.y+math.Sin(p.dir)*15))
	imd.Line(1)

	imd.Draw(t)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Nada",
		Bounds: pixel.R(0, 0, windowWidth, windowHeight),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	last := time.Now()

	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()
		update(win, dt)

		win.Clear(colornames.Skyblue)
		drawMinimap(win)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
