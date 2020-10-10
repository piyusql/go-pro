package main

import (
	"math/rand"
	"time"
)

type Boid struct {
	position vector2D
	velocity vector2D
	id       int
}

func (b *Boid) moveOne() {
	b.position = b.position.Add(b.velocity)
	next := b.position
	if next.x > screenWidth || next.x < 0 {
		b.velocity = vector2D{x: -1 * b.velocity.x, y: b.velocity.y}
	}
	if next.y > screenHeight || next.y < 0 {
		b.velocity = vector2D{x: b.velocity.x, y: -1 * b.velocity.y}
	}
}
func (b *Boid) start() {
	for {
		b.moveOne()
		time.Sleep(10 * time.Millisecond)
	}
}

func createBoid(bid int) {
	b := Boid{
		position: vector2D{x: rand.Float64() * screenWidth, y: rand.Float64() * screenHeight},
		velocity: vector2D{x: (rand.Float64() * 2) - 1.0, y: (rand.Float64()*2 + 1.0)},
		id:       bid,
	}
	boids[bid] = &b
	go b.start()
}
