package main

import (
	"fmt"
	"log"
	"sort"
	"time"
)

const (
	MaxGenerations int     = 1000
	PopulationSize int     = 100
	MutationRate   float64 = 0.0001
	CrossoverRate  float64 = 0.9
	ElitismCount   int     = 10
	TournamentSize int     = 20
)

var PredefinedCities []City = []City{{60, 94}, {66, 43}, {42, 68}, {6, 15}, {9, 30}, {51, 81}, {21, 38}, {31, 46}, {28, 29}, {67, 21}, {20, 36}, {57, 86}, {29, 29}, {75, 20}, {86, 69}, {52, 2}, {15, 60}, {97, 7}, {59, 5}, {69, 30}, {17, 54}, {54, 27}, {42, 53}, {25, 28}, {78, 36}, {88, 29}, {89, 9}, {97, 7}, {22, 68}, {24, 31}, {93, 74}, {80, 73}, {18, 42}, {89, 68}, {97, 92}, {9, 49}, {92, 95}, {34, 69}, {71, 56}, {64, 55}, {75, 40}, {13, 98}, {89, 32}, {72, 64}, {8, 66}, {62, 36}, {23, 53}, {18, 23}, {62, 12}, {28, 41}, {43, 62}, {55, 62}, {72, 83}, {0, 73}, {39, 49}, {60, 40}, {2, 0}, {0, 91}, {58, 55}, {81, 87}, {45, 60}, {2, 84}, {24, 64}, {24, 17}, {59, 81}, {69, 3}, {53, 97}, {75, 29}, {75, 15}, {35, 83}, {23, 62}, {49, 8}, {2, 39}, {58, 92}, {57, 58}, {41, 55}, {49, 95}, {79, 10}, {78, 39}, {13, 19}, {73, 65}, {9, 52}, {9, 15}, {7, 31}, {15, 13}, {32, 53}, {57, 51}, {68, 65}, {52, 65}, {71, 63}, {1, 3}, {9, 36}, {82, 34}, {34, 25}, {21, 55}, {40, 50}, {16, 33}, {82, 70}, {5, 99}, {41, 11}}

func startGenerations() {
	start := time.Now()
	// Create cities
	numCities := 100
	cities := make([]City, numCities)
	//var rand = new Random()
	// Loop to create random cities
	/*for cityIndex := 0; cityIndex < numCities; cityIndex++ {
		// Generate x,y position
		xPos := int(rand.Float64() * 100)
		yPos := int(rand.Float64() * 100)

		// Add city
		cities[cityIndex] = City{xPos, yPos}
	}*/
	//fmt.Println(cities)
	cities = PredefinedCities

	// Initial GA
	ga := GeneticAlgorithm{
		PopulationSize,
		MutationRate,
		CrossoverRate,
		ElitismCount,
		TournamentSize,
	}

	// Initialize population
	population := ga.InitPopulation(len(cities))

	// Evaluate population
	ga.EvalPopulation(&population, cities)

	// my sort
	sort.Sort(ByFitness(population.Population))
	startRoute := createRoute(population.GetFittest(0), cities)
	fmt.Printf("Start Distance: %f\n", startRoute.GetDistance())

	// Keep track of current generation
	generation := 1
	// Start evolution loop
	var route Route
	for ga.IsTerminationConditionMet(generation, MaxGenerations) == false {
		// my sort
		//population.Sort()
		// Print fittest individual from population
		route := createRoute(population.GetFittest(0), cities)
		fmt.Printf("G %d Best distance: %f\n", generation, route.GetDistance())

		// Apply crossover
		population = ga.CrossoverPopulation(population)

		// Apply mutation
		population = ga.MutatePopulation(population)

		// Evaluate population
		ga.EvalPopulation(&population, cities)

		// Increment the current generation
		generation++
	}

	// my sort
	sort.Sort(ByFitness(population.Population))
	fmt.Printf("Stopped after %d generations.\n", generation)
	route = createRoute(population.GetFittest(0), cities)
	fmt.Printf("Best distance: %f\n", route.GetDistance())

	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)
	//Console.ReadKey()
}

func main() {
	startGenerations()
}
