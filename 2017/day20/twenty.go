package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Particle struct {
	id           int
	position     Property
	velocity     Property
	acceleration Property
	distance     int
	destroyed    bool
}

type Property struct {
	x int
	y int
	z int
}

func main() {
	sim1 := createParticles()
	sim2 := createParticles()
	simulate(sim1)
	simulateCollisions(sim2)
}

func simulate(particles []*Particle) {
	notChanged := 0
	lastClosest := tick(particles)
	for {
		// Get the closest particle
		closest := tick(particles)
		// If the closest particle did not change, increase the occurrence counter
		if closest == lastClosest {
			notChanged++
		}
		lastClosest = closest
		// If a particle has been the closest for 5000 iterations, we can most likely assume its the closest forever.
		// Might need some tuning.
		if notChanged == 5000 {
			fmt.Println("Closest: ", closest)
			break
		}
	}
}

func simulateCollisions(particles []*Particle) {
	notChanged := 0
	collisions := 0
	for {
		// Get the closest particle
		collisions = tickCollisions(particles)
		if collisions == 0 {
			notChanged++
		} else {
			notChanged = 0
		}
		// If for 5000 consecutive iterations there have not been collisions, we can assume there wont be any more.
		if notChanged == 100 {
			fmt.Println("Number of particles left: ", particlesLeft(particles))
			break
		}
	}
}

func tick(particles []*Particle) int {
	closest := 0 // Particle ID closest to the 0,0,0
	minDist := math.MaxInt32
	for _, p := range particles {
		// Change position and velocity
		p.velocity.x += p.acceleration.x
		p.velocity.y += p.acceleration.y
		p.velocity.z += p.acceleration.z
		p.position.x += p.velocity.x
		p.position.y += p.velocity.y
		p.position.z += p.velocity.z
		// Calculate its Manhattan distance to 0
		p.distance = int(math.Abs(float64(p.position.x)) + math.Abs(float64(p.position.y)) + math.Abs(float64(p.position.z)))
		// Save the closest
		if p.distance < minDist {
			minDist = p.distance
			closest = p.id
		}
	}
	return closest
}

func tickCollisions(particles []*Particle) int {
	// Change position and velocity
	for _, p := range particles {
		if !p.destroyed {
			p.velocity.x += p.acceleration.x
			p.velocity.y += p.acceleration.y
			p.velocity.z += p.acceleration.z
			p.position.x += p.velocity.x
			p.position.y += p.velocity.y
			p.position.z += p.velocity.z
		}
	}
	// Check Collisions
	totalCollisions := 0
	for _, p1 := range particles {
		// Collisions with this particle
		collisions := 0
		for _, p2 := range particles {
			// Dont check the same particle, nor if any of them has already been destroyed
			if p1.id != p2.id && !p1.destroyed && !p2.destroyed {
				if (p1.position.x == p2.position.x) && (p1.position.y == p2.position.y) && (p1.position.z == p2.position.z) {
					// Collision
					collisions++
					p2.destroyed = true
				}
			}
		}
		// If the particle we checked against collided with something. Destroy it.
		// If we had set it to destroyed earlier, the if check would fail too early, upon the first collision.
		if collisions > 0 {
			p1.destroyed = true
		}
		totalCollisions += collisions
		collisions = 0
	}
	return totalCollisions
}

func createParticles() []*Particle {
	particles := []*Particle{}
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		prpt := strings.Fields(line)
		p, _ := strings.CutSuffix(prpt[0][3:], ">,")
		v, _ := strings.CutSuffix(prpt[1][3:], ">,")
		a, _ := strings.CutSuffix(prpt[2][3:], ">")
		pos := strings.Split(p, ",")
		vel := strings.Split(v, ",")
		accel := strings.Split(a, ",")

		px, _ := strconv.Atoi(pos[0])
		py, _ := strconv.Atoi(pos[1])
		pz, _ := strconv.Atoi(pos[2])

		vx, _ := strconv.Atoi(vel[0])
		vy, _ := strconv.Atoi(vel[1])
		vz, _ := strconv.Atoi(vel[2])

		ax, _ := strconv.Atoi(accel[0])
		ay, _ := strconv.Atoi(accel[1])
		az, _ := strconv.Atoi(accel[2])

		particles = append(particles,
			&Particle{i, Property{px, py, pz},
				Property{vx, vy, vz},
				Property{ax, ay, az}, 0, false})
		i++
	}
	return particles
}

func particlesLeft(particles []*Particle) int {
	count := 0
	for i := range particles {
		if !particles[i].destroyed {
			count++
		}
	}
	return count
}
