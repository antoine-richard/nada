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
	rotationSpeed = 3
	walkSpeed     = 30
)

type player struct {
	x, y float64
	dir  float64
}

var p player

var walls = []pixel.Vec{
	pixel.V(400, 150),
	pixel.V(500, 250),
	pixel.V(350, 400),
	pixel.V(250, 300),
	pixel.V(200, 150),
}

var projectedWalls = []pixel.Vec{
	pixel.ZV,
	pixel.ZV,
	pixel.ZV,
	pixel.ZV,
	pixel.ZV,
}

func processInput(win *pixelgl.Window, dt float64) {
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

func updateWorld(win *pixelgl.Window) {
	cs := math.Cos(math.Pi/2 - p.dir)
	sn := math.Sin(math.Pi/2 - p.dir)
	for i := 0; i < len(walls); i++ {
		projectedWalls[i].X = (walls[i].X-p.x)*cs - (walls[i].Y-p.y)*sn + win.Bounds().Center().X
		projectedWalls[i].Y = (walls[i].X-p.x)*sn + (walls[i].Y-p.y)*cs + win.Bounds().Center().Y
	}
}

func drawMinimap(win *pixelgl.Window, minimap *pixelgl.Canvas) {
	minimap.Clear(colornames.Skyblue)
	imd := imdraw.New(nil)

	// player position
	imd.Color = colornames.Navy
	imd.Push(pixel.V(p.x, p.y))
	imd.Circle(7, 0)

	// player direction
	imd.Push(pixel.V(p.x, p.y), pixel.V(p.x+math.Cos(p.dir)*30, p.y+math.Sin(p.dir)*30))
	imd.Line(4)

	// world
	imd.Push(walls...)
	imd.Polygon(4)

	imd.Draw(minimap)
	minimap.Draw(
		win,
		pixel.IM.
			Scaled(minimap.Bounds().Center(), 0.25).
			Moved(pixel.V(-minimap.Bounds().W()/4, minimap.Bounds().H()/2)),
	)
}

func drawWorld(win *pixelgl.Window) {
	imd := imdraw.New(nil)

	// fixed player position
	imd.Color = colornames.Red
	imd.Push(win.Bounds().Center())
	imd.Circle(3, 0)

	// fixed player direction
	imd.Push(win.Bounds().Center(), win.Bounds().Center().Add(pixel.V(0, 15)))
	imd.Line(1)

	// rotated world
	imd.Color = colornames.Green
	imd.Push(projectedWalls...)
	imd.Polygon(1)

	imd.Draw(win)
}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Nada",
		Bounds: pixel.R(0, 0, 640, 480),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	p = player{
		x:   win.Bounds().W() / 2,
		y:   win.Bounds().H() / 2,
		dir: math.Pi * 3 / 4,
	}

	minimap := pixelgl.NewCanvas(pixel.R(0, 0, win.Bounds().W(), win.Bounds().H()))
	last := time.Now()

	for !win.Closed() {
		dt := time.Since(last).Seconds()
		last = time.Now()

		processInput(win, dt)

		updateWorld(win)

		win.Clear(colornames.Black)
		drawWorld(win)
		drawMinimap(win, minimap)

		win.Update()
	}
}

func main() {
	pixelgl.Run(run)
}
