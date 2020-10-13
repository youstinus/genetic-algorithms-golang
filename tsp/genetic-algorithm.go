package main

import (
	"math"
	"math/rand"
	"sort"
)

// GeneticAlgorithm re..
type GeneticAlgorithm struct {
	PopulationSize int
	MutationRate   float64
	CrossoverRate  float64
	ElitismCount   int
	TournamentSize int
}

func createGeneticAlgorithm() {

}

/**
 * Initialize population
 *
 * @param chromosomeLength The length of the individuals chromosome
 * @return population The initial population generated
 */
func (g *GeneticAlgorithm) InitPopulation(chromosomeLength int) Population {
	// Initialize population
	return createPopulation(g.PopulationSize, chromosomeLength)
}

/**
 * Check if population has met termination condition -- this termination
 * condition is a simple one simply check if we've exceeded the allowed
 * number of generations.
 *
 * @param generationsCount
 *            Number of generations passed
 * @param maxGenerations
 *            Number of generations to terminate after
 * @return boolean True if termination condition met, otherwise, false
 */
func (g *GeneticAlgorithm) IsTerminationConditionMet(generationsCount int, maxGenerations int) bool {
	return generationsCount > maxGenerations
}

/**
 * Calculate individual's fitness value
 *
 * Fitness, in this problem, is inversely proportional to the route's total
 * distance. The total distance is calculated by the Route class.
 *
 * @param individual
 *            the individual to evaluate
 * @param cities
 *            the cities being referenced
 * @return double The fitness value for individual
 */
func (g *GeneticAlgorithm) CalcFitness(individual *Individual, cities []City) float64 {
	// Get fitness
	var route = createRoute(*individual, cities)
	var fitness = 1 / route.GetDistance()

	// Store fitness
	individual.Fitness = fitness

	return fitness
}

/**
 * Evaluate population -- basically run calcFitness on each individual.
 *
 * @param population the population to evaluate
 * @param cities the cities being referenced
 */
func (g *GeneticAlgorithm) EvalPopulation(population *Population, cities []City) {
	populationFitness := 0.0

	for i := range population.Population {
		populationFitness += g.CalcFitness(&population.Population[i], cities)
	}
	//.Sum(individual => CalcFitness(individual, cities))

	// Loop over population evaluating individuals and summing population fitness

	avgFitness := populationFitness / float64(len(population.Population))
	population.PopulationFitness = avgFitness
}

/**
 * Selects parent for crossover using tournament selection
 *
 * Tournament selection was introduced in Chapter 3
 *
 * @param population
 *
 * @return The individual selected as a parent
 */
func (g *GeneticAlgorithm) SelectParent(population Population) Individual {
	// Create tournament
	tournament := CreateTournamentPopulation(g.TournamentSize)

	// Add random individuals to the tournament
	population.shuffle()
	for i := 0; i < g.TournamentSize; i++ {
		var tournamentIndividual = population.Population[i]
		tournament.Population[i] = tournamentIndividual
	}

	// Return the best
	//tournament.Sort()
	return tournament.GetFittest(0)
}

/**
 * Ordered crossover mutation
 *
 * Chromosomes in the TSP require that each city is visited exactly once.
 * Uniform crossover can break the chromosome by accidentally selecting a
 * city that has already been visited from a parent this would lead to one
 * city being visited twice and another city being skipped altogether.
 *
 * Additionally, uniform or random crossover doesn't really preserve the
 * most important aspect of the genetic information: the specific order of a
 * group of cities.
 *
 * We need a more clever crossover algorithm here. What we can do is choose
 * two pivot points, add chromosomes from one parent for one of the ranges,
 * and then only add not-yet-represented cities to the second range. This
 * ensures that no cities are skipped or visited twice, while also
 * preserving ordered batches of cities.
 *
 * @param population
 * @return The new population
 */
func (g *GeneticAlgorithm) CrossoverPopulation(population Population) Population {
	// Create new population
	var newPopulation = Population{
		Population: make([]Individual, len(population.Population)),
	} //createPopulation(len(population.Population))

	// my sort
	sort.Sort(ByFitness(population.Population))
	// Loop over current population by fitness
	for populationIndex := 0; populationIndex < len(population.Population); populationIndex++ {
		// Get parent1
		var parent1 = population.GetFittest(populationIndex)

		// Apply crossover to this individual?
		if g.CrossoverRate > rand.Float64() && populationIndex >= g.ElitismCount {
			// Find parent2 with tournament selection
			var parent2 = g.SelectParent(population)

			// Create blank offspring chromosome
			var offspringChromosome = make([]int, len(parent1.Chromosome))
			offspringChromosome = repeat(-1, len(offspringChromosome)) //Enumerable.Repeat(-1, offspringChromosome.Length).ToArray()
			//Arrays.fill(offspringChromosome, -1)
			var offspring = Individual{
				Chromosome: offspringChromosome,
			}

			// Get subset of parent chromosomes
			var substrPos1 = (int)(rand.Float64() * float64(len(parent1.Chromosome)))
			var substrPos2 = (int)(rand.Float64() * float64(len(parent1.Chromosome)))

			// make the smaller the start and the larger the end
			var startSubstr = int(math.Min(float64(substrPos1), float64(substrPos2)))
			var endSubstr = int(math.Max(float64(substrPos1), float64(substrPos2)))

			// Loop and add the sub tour from parent1 to our child
			for i := startSubstr; i < endSubstr; i++ {
				offspring.Chromosome[i] = parent1.Chromosome[i]
			}

			// Loop through parent2's city tour
			for i := 0; i < len(parent2.Chromosome); i++ {
				var parent2Gene = i + endSubstr
				if parent2Gene >= len(parent2.Chromosome) {
					parent2Gene -= len(parent2.Chromosome)
				}

				// If offspring doesn't have the city add it
				if offspring.ContainsGene(parent2.Chromosome[parent2Gene]) == false {
					// Loop to find a spare position in the child's tour
					for ii := 0; ii < len(offspring.Chromosome); ii++ {
						// Spare position found, add city
						if offspring.Chromosome[ii] == -1 {
							offspring.Chromosome[ii] = parent2.Chromosome[parent2Gene]
							break
						}
					}
				}
			}

			// Add child
			newPopulation.Population[populationIndex] = offspring
		} else {
			// Add individual to new population without applying crossover
			newPopulation.Population[populationIndex] = parent1
		}
	}

	return newPopulation
}

/**
 * Apply mutation to population
 *
 * Because the traveling salesman problem must visit each city only once,
 * this form of mutation will randomly swap two genes instead of
 * bit-flipping a gene like in earlier examples.
 *
 * @param population
 *            The population to apply mutation to
 * @return The mutated population
 */
func (g *GeneticAlgorithm) MutatePopulation(population Population) Population {
	// Initialize new population
	//var newPopulation = createPopulation(g.PopulationSize)
	// Create new population
	var newPopulation = Population{
		Population: make([]Individual, len(population.Population)),
	} //createPopulation(len(population.Population))

	// my sort
	sort.Sort(ByFitness(population.Population))
	// Loop over current population by fitness
	for populationIndex := 0; populationIndex < len(population.Population); populationIndex++ {
		var individual = population.GetFittest(populationIndex)

		// Skip mutation if this is an elite individual
		if populationIndex >= g.ElitismCount {
			// System.out.println("Mutating population member "+populationIndex)
			// Loop over individual's genes
			for geneIndex := 0; geneIndex < len(individual.Chromosome); geneIndex++ {
				// System.out.println("\tGene index "+geneIndex)
				// Does this gene need mutation?
				if g.MutationRate > rand.Float64() {
					// Get new gene position
					var newGenePos = int(rand.Float64() * float64(len(individual.Chromosome)))
					// Get genes to swap
					var gene1 = individual.Chromosome[newGenePos]
					var gene2 = individual.Chromosome[geneIndex]
					// Swap genes
					individual.Chromosome[geneIndex] = gene1
					individual.Chromosome[newGenePos] = gene2
				}
			}
		}

		// Add individual to population
		newPopulation.Population[populationIndex] = individual
	}

	// Return mutated population
	return newPopulation
}

func repeat(element, count int) []int {
	elems := make([]int, count)
	for i := 0; i < count; i++ {
		elems[i] = element
	}
	return elems
}
