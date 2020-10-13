package main

// Route re..
type Route struct {
	Route    []City
	Distance float64
}

// CreateRoute from cities and chromosome arangeament
func createRoute(individual Individual, cities []City) Route {
	// Get individual's chromosome
	var chromosome []int = individual.Chromosome
	// Create route
	route := make([]City, len(cities))
	for geneIndex := 0; geneIndex < len(chromosome); geneIndex++ {
		route[geneIndex] = cities[chromosome[geneIndex]]
	}
	return Route{route, 0.0}
}

// GetDistance returns route distance
func (r *Route) GetDistance() float64 {
	if r.Distance > 0 {
		return r.Distance
	}
	n := len(r.Route)
	// Loop over cities in route and calculate route distance
	var totalDistance float64 = 0
	for cityIndex := 0; cityIndex+1 < n; cityIndex++ {
		totalDistance += r.Route[cityIndex].DistanceFrom(r.Route[cityIndex+1])
	}

	totalDistance += r.Route[n-1].DistanceFrom(r.Route[0])
	r.Distance = totalDistance

	return totalDistance
}
