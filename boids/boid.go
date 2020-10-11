package boids

import (
	"math"
	"math/rand"
	"time"
)

type Boid struct {
	position vector2D
	velocity vector2D
	id       int
}

func (b *Boid) createAcceleration() vector2D {
	rWLock.RLock()
	avgVelocity, avgPosition, seperation := vector2D{x: 0, y: 0}, vector2D{x: 0, y: 0}, vector2D{x: 0, y: 0}
	count := 0.0
	lower, upper := b.position.AddV(-viewRadius), b.position.AddV(viewRadius)
	for i := int(math.Max(0, lower.x)); i <= int(math.Min(upper.x, screenWidth)); i++ {
		for j := int(math.Max(0, lower.y)); j <= int(math.Min(upper.y, screenHeight)); j++ {
			if otherId := boidMap[i][j]; otherId != -1 && otherId != b.id {
				if dist := b.position.Distance(boids[otherId].position); dist < viewRadius {
					avgVelocity = avgVelocity.Add(boids[otherId].velocity)
					avgPosition = avgPosition.Add(boids[otherId].position)
					seperation = seperation.Add(b.position.Subtract(boids[otherId].position).DivisionV(dist))
					count++
				}
			}
		}
	}
	rWLock.RUnlock()
	accel := vector2D{b.borderBounce(b.position.x, screenWidth), b.borderBounce(b.position.y, screenHeight)}
	if count > 0 {
		avgVelocity, avgPosition = avgVelocity.DivisionV(count), avgPosition.DivisionV(count)
		accelAlignment := avgVelocity.Subtract(b.velocity)
		accelCohesion := avgPosition.Subtract(b.position)
		accelSeperation := seperation
		accel = accel.Add(accelAlignment).Add(accelCohesion).Add(accelSeperation)
	}
	return accel.MultiplyV(slowFactor)
}

func (b *Boid) borderBounce(pos, maxScreen float64) float64 {
	if pos < viewRadius {
		return 1 / pos
	} else if pos > maxScreen-viewRadius {
		return 1 / (pos - maxScreen)
	}
	return 0
}

func (b *Boid) moveOne() {
	accel := b.createAcceleration()
	b.velocity = b.velocity.Add(accel).Limit(-1, 1)
	rWLock.Lock()
	boidMap[b.position.X()][b.position.Y()] = -1
	b.position = b.position.Add(b.velocity)
	// edge condition if goying beyond screen
	bouncing := false
	next := b.position
	if next.x > screenWidth || next.x < 0 {
		b.velocity = vector2D{x: -1 * b.velocity.x, y: b.velocity.y}
		bouncing = true
	}
	if next.y > screenHeight || next.y < 0 {
		b.velocity = vector2D{x: b.velocity.x, y: -1 * b.velocity.y}
		bouncing = true
	}
	if bouncing {
		b.position = b.position.Add(b.velocity)
	}
	boidMap[b.position.X()][b.position.Y()] = b.id
	rWLock.Unlock()
}

func (b *Boid) start() {
	for {
		b.moveOne()
		time.Sleep(5 * time.Millisecond)
	}
}

func createBoid(bid int) {
	b := Boid{
		position: vector2D{x: rand.Float64() * screenWidth, y: rand.Float64() * screenHeight},
		velocity: vector2D{x: (rand.Float64() * 2) - 1.0, y: (rand.Float64() * 2) - 1.0},
		id:       bid,
	}
	boids[bid] = &b
	boidMap[b.position.X()][b.position.Y()] = b.id
	go b.start()
}
