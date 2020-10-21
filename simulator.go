package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	// Set up rand
	rand.Seed(time.Now().UnixNano())

	printInterval := 10000000
	gameCount := 0
	totalScore := 0

	for {
		for i := 0; i < printInterval; i++ {
			totalScore += playGame()
		}
		gameCount += printInterval

		avgScore := float64(totalScore) / float64(gameCount)
		fmt.Printf("After %s games the average is %f\n", humanize(gameCount), avgScore)
	}
}

// Return nice string for humans
func humanize(n int) string {
	if n >= 10000000000 {
		return strconv.Itoa(n/1000000000) + " billion"
	}

	if n >= 10000000 {
		return strconv.Itoa(n/1000000) + " million"
	}

	return strconv.Itoa(n)
}

// Play the dart game, returning the score
// It makes it a bit ugly, but using the squared values makes calculations a lot
// more efficient, as it avoids lots of square roots and multiplications.
func playGame() int {
	score := 1
	radiusSquared := 1.0

	// Simulate throws
	for {
		// This should calculate e^(pi/4)
		distanceSquared := randomDistanceSquaredToCentreFromSquare()

		// This should calculate e
		// distanceSquared := randomDistanceSquaredToCentreFromCircle()

		// If we've thrown outside the dartboard
		if distanceSquared > radiusSquared {
			break
		}

		radiusSquared = chordLengthSquared(distanceSquared, radiusSquared)
		score++
	}

	return score
}

// Length of a chord with midpoint a given distance from the centre, in a circle
// of given radius
func chordLengthSquared(distanceSquared, radiusSquared float64) float64 {
	return radiusSquared - distanceSquared
}

// Returns the distance squared to (0, 0) from a random point (uniformly
// distributed by area) with x and y in range (-1, 1)
//
// This is the distrubtion used in the original puzzle
func randomDistanceSquaredToCentreFromSquare() float64 {
	// We can just use random between (0, 1) as the distribution is symmetric
	// about the x and y axes
	x := rand.Float64()
	y := rand.Float64()
	return x*x + y*y
}

// Returns the distance to (0, 0) from a random point (uniformly distributed by
// area) with |p| < 1
func randomDistanceSquaredToCentreFromCircle() float64 {
	// We can just use random between (0, 1) mapped through sqrt (and as we're
	// returning squared just the random distribution itself)
	//
	// The distribution is radially symmetric so we can just return the radius.
	// The circumference of a circle (which is how many points you'd 'fit' at
	// that radius) is 2*pi*radius, so the probability of a uniformly
	// distributed point by area landing at a radius is proportional to that
	// radius. To find cumulative distribution this is integrated to something
	// proportional to radius^2. Then the inverse (the mapping to transform the
	// uniform distribution to the desired one) is sqrt(radius). Then we square
	// it (as we're returning distance squared) so we can just return the
	// uniform distribution.
	return rand.Float64()
}
