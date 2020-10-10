package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth, screenHeight = 640, 360
	boidCount                 = 500
)

var (
	green  = color.RGBA{R: 10, G: 255, B: 50, A: 255}
	red    = color.RGBA{R: 255, G: 10, B: 50, A: 255}
	colors = [2]color.RGBA{green, red}
	boids  [boidCount]*Boid
)

func update(screen *ebiten.Image) error {
	/* creating a bird in the followig position
		*
	   * *
		*
		*
	*/
	if !ebiten.IsDrawingSkipped() {
		for i, boid := range boids {
			screen.Set(int(boid.position.x+1), int(boid.position.y), colors[i%2])
			screen.Set(int(boid.position.x-1), int(boid.position.y), colors[i%2])
			screen.Set(int(boid.position.x), int(boid.position.y-1), colors[i%2])
			screen.Set(int(boid.position.x), int(boid.position.y+1), colors[i%2])
			screen.Set(int(boid.position.x), int(boid.position.y+2), colors[i%2])
		}
	}
	return nil
}

func main() {
	for i := 0; i < boidCount; i++ {
		createBoid(i)
	}
	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Birds waving !!!"); err != nil {
		log.Fatal(err)
	}
}
