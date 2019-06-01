package main

// Population represents ...
type Population struct {
	Population        []Individual
	PopulationFitness float64
}

// creates the initial population
func createPopulation(target []rune) (population []Individual) {
	population = make([]Individual, PopSize)
	for i := 0; i < PopSize; i++ {
		population[i] = createIndividual(target)
	}
	return
}
