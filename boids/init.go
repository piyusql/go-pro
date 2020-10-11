package boids

import (
	"image/color"
	"log"
	"sync"

	"github.com/hajimehoshi/ebiten"
)

const (
	screenWidth, screenHeight = 640, 360
	boidCount                 = 1024
	viewRadius                = 13
	slowFactor                = 0.125
)

var (
	skycolour  = color.RGBA{R: 185, G: 235, B: 255, A: 255}
	birdcolour = color.RGBA{R: 1, G: 4, B: 16, A: 255}
	boids      [boidCount]*Boid
	// this will store the boid id at the screen position
	boidMap [screenWidth + 1][screenHeight + 1]int
	rWLock  = sync.RWMutex{}
)

func update(screen *ebiten.Image) error {
	/* creating a bird in the followig pixel position
		*
	   * *
		*
		*
	*/
	if !ebiten.IsDrawingSkipped() {
		screen.Fill(skycolour)
		for _, boid := range boids {
			screen.Set(int(boid.position.x-1), int(boid.position.y), birdcolour)
			screen.Set(int(boid.position.x+1), int(boid.position.y), birdcolour)
			screen.Set(int(boid.position.x), int(boid.position.y-1), birdcolour)
			screen.Set(int(boid.position.x), int(boid.position.y+1), birdcolour)
			screen.Set(int(boid.position.x), int(boid.position.y+2), birdcolour)
		}
	}
	return nil
}

func Init() {
	for i, row := range boidMap {
		for j := range row {
			boidMap[i][j] = -1
		}
	}
	for i := 0; i < boidCount; i++ {
		createBoid(i)
	}
	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Flocking birds, https://github.com/piyusgupta/go-pro !!!"); err != nil {
		log.Fatal(err)
	}
}
