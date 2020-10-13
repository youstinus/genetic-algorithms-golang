package main

import "math"

// City re..
type City struct {
	X int
	Y int
}

func createCity() {

}

// DistanceFrom returns distance between cities
func (c *City) DistanceFrom(city City) float64 {
	// Give difference in x,y
	var deltaXSq = math.Pow(float64(city.X-c.X), 2.0)
	var deltaYSq = math.Pow(float64(city.Y-c.Y), 2.0)

	// Calculate shortest path
	var distance = math.Sqrt(math.Abs(deltaXSq + deltaYSq))
	return distance
}
