package main

import (
	"math"
	"math/rand"
	"time"
)

// GeneratePlayerLocations generates player's initial locations in the world
func GeneratePlayerLocations(numPlayers int, worldRadius float64) (locs []Location) {

	rand.Seed(time.Now().Unix())

	locs = make([]Location, numPlayers)

	angle := rand.Float64() * 2 * math.Pi
	avgSpace := (2 * math.Pi) / float64(numPlayers)
	innerMargin := worldRadius * 0.4
	outerMargin := worldRadius * 0.8
	spread := math.Pi / 12

	for i := 0; i < numPlayers; i++ {
		angle += avgSpace + rand.Float64()*(2*spread) - spread
		dist := rand.Float64()*(outerMargin-innerMargin) + innerMargin

		x := math.Cos(angle) * dist
		y := math.Sin(angle) * dist

		locs[i] = Location{x, y}
	}
	return
}
