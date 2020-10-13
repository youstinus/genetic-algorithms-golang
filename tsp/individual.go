package main

import "math/rand"

// Individual re..
type Individual struct {
	Chromosome []int
	Fitness    float64
}

func createIndividual(chromosomeLength int) (i Individual) {
	// Create random individual
	individual := make([]int, chromosomeLength)

	/**
	 * In this case, we can no longer simply pick 0s and 1s -- we need to
	 * use every city index available. We also don't need to randomize or
	 * shuffle this chromosome, as crossover and mutation will ultimately
	 * take care of that for us.
	 */
	for gene := 0; gene < chromosomeLength; gene++ {
		individual[gene] = gene
	}
	individual = shuffle(individual)

	i.Chromosome = individual
	return i
}

/**
 * Search for a specific integer gene in this individual.
 *
 * For instance, in a Traveling Salesman Problem where cities are encoded as
 * integers with the range, say, 0-99, this method will check to see if the
 * city "42" exists.
 *
 * @param gene
 * @return
 */
func (i *Individual) ContainsGene(gene int) bool {
	for ii := range i.Chromosome {
		if i.Chromosome[ii] == gene {
			return true
		}
	}
	return false //_chromosome.Any(t => t == gene);
}

// shuffle s
/**
 * Shuffles the population in-place
 *
 * @param void
 * @return void
 */
func shuffle(numbers []int) []int {
	for i := len(numbers) - 1; i > 0; i-- {
		index := rand.Intn(i + 1)
		/*a := p.Population[index]
		p.Population[index] = p.Population[i]
		p.Population[i] = a*/ // instead of this
		numbers[index], numbers[i] = numbers[i], numbers[index]
	}
	return numbers
}
