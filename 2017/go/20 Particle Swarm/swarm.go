package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Particle struct {
	ID              int
	AbsAcceleration float64
	Position        Position
	Velocity        Position
	Acceleration    Position
}
type Position struct {
	X float64
	Y float64
	Z float64
}
type ByVelocity []Particle

func (a ByVelocity) Len() int           { return len(a) }
func (a ByVelocity) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByVelocity) Less(i, j int) bool { return a[i].AbsAcceleration < a[j].AbsAcceleration }

func main() {
	lines := fileToLines("input.txt")
	particles := []Particle{}

	for id, line := range lines {
		p := parseLine(id, line)
		particles = append(particles, p)
	}
	// sort.Sort(ByVelocity(particles))

	// fmt.Printf("--------------------------------")
	// fmt.Printf("Part one\n %+v \n %+v \n", particles[0], particles[len(particles)-1])
	deleted := map[int]int{}

	for t := 1; t < 100; t++ {
		for ia, a := range particles {

			for ib, b := range particles {
				if a.ID == b.ID {
					continue
				} else {
				}
				if IsCollition(a, b) {
					deleted[a.ID] = ia
					deleted[b.ID] = ib
				}

			}
		}

		for id := range deleted {
			idx := findParticleIndex(id, particles)
			particles = append(particles[:idx], particles[idx+1:]...)
		}
		deleted = map[int]int{}

		//Move time
		for i := 0; i < len(particles); i++ {
			particles[i] = particles[i].Move()
		}

		fmt.Printf("Count %d %d\n", len(particles), t)
	}

	sort.Sort(ByVelocity(particles))
	fmt.Printf("--------------------------------")
	fmt.Printf("Part two\n %d \n ", len(particles))

}

func findParticleIndex(id int, particles []Particle) int {
	for i, p := range particles {
		if p.ID == id {
			return i
		}
	}
	return -1
}
func IsCollition(a, b Particle) bool {
	return a.Position.X == b.Position.X &&
		a.Position.Y == b.Position.Y &&
		a.Position.Z == b.Position.Z

}
func (p Particle) Move() Particle {
	p.Velocity.X += p.Acceleration.X
	p.Velocity.Y += p.Acceleration.Y
	p.Velocity.Z += p.Acceleration.Z

	p.Position.X += p.Velocity.X
	p.Position.Y += p.Velocity.Y
	p.Position.Z += p.Velocity.Z

	return p
}

func willCollide(a, b Particle) bool {
	tx := 0.
	var pa, va, pb, vb float64

	found := false

	pa = a.Position.X
	va = a.Velocity.X
	pb = b.Position.X
	vb = b.Velocity.X
	for i := 0; i < 1000; i++ {

		if pa == pb {
			found = true
			break
		}
		tx++
		va += a.Acceleration.X
		pa += va

		vb += b.Acceleration.X
		pb += vb
	}
	if !found {
		return false
	}

	return true
}

func parseLine(id int, line string) Particle {
	ret := Particle{ID: id}

	parts := strings.Split(line, ",")
	xs := strings.Replace(parts[6], " a=<", "", -1)
	xs = strings.TrimSpace(xs)
	ys := parts[7]
	zs := strings.TrimSuffix(parts[8], ">")

	ret.Acceleration.X, _ = strconv.ParseFloat(xs, 64)
	ret.Acceleration.Y, _ = strconv.ParseFloat(ys, 64)
	ret.Acceleration.Z, _ = strconv.ParseFloat(zs, 64)
	ret.AbsAcceleration = math.Abs(ret.Acceleration.X) + math.Abs(ret.Acceleration.Y) + math.Abs(ret.Acceleration.Z)

	xs = strings.Replace(parts[0], "p=<", "", -1)
	xs = strings.TrimSpace(xs)
	ys = parts[1]
	zs = strings.TrimSuffix(parts[2], ">")

	ret.Position.X, _ = strconv.ParseFloat(xs, 64)
	ret.Position.Y, _ = strconv.ParseFloat(ys, 64)
	ret.Position.Z, _ = strconv.ParseFloat(zs, 64)

	xs = strings.Replace(parts[3], " v=<", "", -1)
	xs = strings.TrimSpace(xs)
	ys = parts[4]
	zs = strings.TrimSuffix(parts[5], ">")

	ret.Velocity.X, _ = strconv.ParseFloat(xs, 64)
	ret.Velocity.Y, _ = strconv.ParseFloat(ys, 64)
	ret.Velocity.Z, _ = strconv.ParseFloat(zs, 64)

	return ret
}

func possibleAxisCollide(a, b Particle) bool {
	da := math.Abs(a.Position.X - b.Position.X)
	db := math.Abs((a.Position.X + a.Acceleration.X) - (b.Position.X + b.Acceleration.X))
	if da < db {
		return false
	}

	da = math.Abs(a.Position.Y - b.Position.Y)
	db = math.Abs((a.Position.Y + a.Acceleration.Y) - (b.Position.Y + b.Acceleration.Y))
	if da < db {
		return false
	}

	da = math.Abs(a.Position.Z - b.Position.Z)
	db = math.Abs((a.Position.Z + a.Acceleration.Z) - (b.Position.Z + b.Acceleration.Z))
	if da < db {
		return false
	}

	return true
}

func getVelocity(line string) float64 {
	parts := strings.Split(line, ",")
	vxs := strings.TrimPrefix(parts[6], " a=<")
	vys := parts[7]
	vzs := strings.TrimSuffix(parts[8], ">")

	vx, _ := strconv.ParseFloat(vxs, 64)
	vy, _ := strconv.ParseFloat(vys, 64)
	vz, _ := strconv.ParseFloat(vzs, 64)
	return math.Abs(vx) + math.Abs(vy) + math.Abs(vz)
}
func fileToLines(fileName string) []string {
	ret := make([]string, 0, 100)

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			ret = append(ret, line)
		}
	}

	return ret
}
