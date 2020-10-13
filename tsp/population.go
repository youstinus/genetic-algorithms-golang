package main

import (
	"math/rand"
	"sort"
)

// Population re..
type Population struct {
	Population        []Individual
	PopulationFitness float64
}

// ByFitness implements sort.Interface for []Individual based on
// the Fitness field.
type ByFitness []Individual

func (a ByFitness) Len() int           { return len(a) }
func (a ByFitness) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFitness) Less(i, j int) bool { return a[i].Fitness > a[j].Fitness } // for desc order

/**
 * Initializes population of individuals
 *
 * @param populationSize
 *            The size of the population
 * @param chromosomeLength
 *            The length of the individuals chromosome
 */
func createPopulation(populationSize, chromosomeLength int) (p Population) {
	// Initial population
	p.Population = make([]Individual, populationSize)

	// Loop over population size
	for individualCount := 0; individualCount < populationSize; individualCount++ {
		// Create individual
		var individual = createIndividual(chromosomeLength)
		// Add individual to population
		p.Population[individualCount] = individual
	}
	return p
}

func CreateTournamentPopulation(tSize int) Population {
	return Population{make([]Individual, tSize), 0.0}
}

// GetFittest s
/*
 * Find fittest individual in the population
 *
 * @param offset
 * @return individual Fittest individual at offset
 */
func (p *Population) GetFittest(offset int) Individual {
	sort.Sort(ByFitness(p.Population))
	return p.Population[offset]
}

// shuffle s
/**
 * Shuffles the population in-place
 *
 * @param void
 * @return void
 */
func (p *Population) shuffle() {
	for i := len(p.Population) - 1; i > 0; i-- {
		index := rand.Intn(i + 1)
		/*a := p.Population[index]
		p.Population[index] = p.Population[i]
		p.Population[i] = a*/ // instead of this
		p.Population[index], p.Population[i] = p.Population[i], p.Population[index]
	}
}

// sortPopulation s
func (p *Population) sortPopulation() {
	sort.Sort(ByFitness(p.Population))
}
