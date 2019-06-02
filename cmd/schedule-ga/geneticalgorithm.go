package main

import (
	"math"
	"math/rand"
)

// GeneticAlgorithm represents
type GeneticAlgorithm struct {
	PopulationSize int
	MutationRate   float64
	CrossoverRate  float64
	ElitismCount   int
	TournamentSize int
}

func createGeneticAlgorithm(populationSize int, mutationRate float64, crossoverRate float64, elitismCount int, tournamentSize int) (geneticAlgorithm GeneticAlgorithm) {
	geneticAlgorithm = GeneticAlgorithm{
		PopulationSize: populationSize,
		MutationRate:   mutationRate,
		CrossoverRate:  crossoverRate,
		ElitismCount:   elitismCount,
		TournamentSize: tournamentSize,
	}
	return
}

/**
 * Initialize population
 *
 * @param chromosomeLength
 *            The length of the individuals chromosome
 * @return population The initial population generated
 */
func (ga *GeneticAlgorithm) InitPopulation(timeTable TimeTable) (population Population) {
	// Initialize population
	population = Population{
		Population:        make([]Individual, ga.PopulationSize),
		PopulationFitness: 0,
	}
	return population
}

/**
 * Check if population has met termination condition
 *
 * @param generationsCount
 *            Number of generations passed
 * @param maxGenerations
 *            Number of generations to terminate after
 * @return boolean True if termination condition met, otherwise, false
 */
func (ga *GeneticAlgorithm) IsTerminationConditionMet1(generationsCount int, maxGenerations int) bool {
	return generationsCount > maxGenerations
}

/**
 * Check if population has met termination condition
 *
 * @param population
 * @return boolean True if termination condition met, otherwise, false
 */
func (ga *GeneticAlgorithm) IsTerminationConditionMet2(population Population) bool {
	return math.Abs(population.GetFittest(0).Fitness-1.0) < 0.00001
}

/**
 * Calculate individual's fitness value
 *
 * @param individual
 * @param timetable
 * @return fitness
 */
func CalcFitness(individual Individual, timeTable TimeTable) float64 {

	// Create new timetable object to use -- cloned from an existing timetable
	threadTimeTable := timeTable
	threadTimeTable.createClasses(individual)

	// Calculate fitness
	clashes := threadTimeTable.calcClashes()
	fitness := 1 / (float64)(clashes+1)

	individual.Fitness = fitness

	return fitness
}

/**
 * Evaluate population
 *
 * @param population
 * @param timetable
 */
func (ga *GeneticAlgorithm) EvalPopulation(population Population, timeTable TimeTable) {
	var populationFitness float64 = 0

	// Loop over population evaluating individuals and summing population
	// fitness
	for _, individual := range population.Population {
		populationFitness += calcFitness(individual, timeTable)
	}

	population.PopulationFitness = populationFitness
}

/**
 * Selects parent for crossover using tournament selection
 *
 * Tournament selection works by choosing N random individuals, and then
 * choosing the best of those.
 *
 * @param population
 * @return The individual selected as a parent
 */
func (ga *GeneticAlgorithm) SelectParent(population Population) Individual {
	// Create tournament
	tournament := Population{
		Population:        make([]Individual, ga.TournamentSize),
		PopulationFitness: 0,
	}

	// Add random individuals to the tournament
	population.shuffle()
	for i := 0; i < ga.TournamentSize; i++ {
		var tournamentIndividual = population.Population[i]
		tournament.Population[i] = tournamentIndividual
	}

	// Return the best
	return tournament.GetFittest(0)
}

/**
 * Apply mutation to population
 *
 * @param population
 * @param timetable
 * @return The mutated population
 */
func (ga *GeneticAlgorithm) MutatePopulation(population Population, timeTable TimeTable) Population {
	// Initialize new population
	newPopulation := Population{
		Population:        make([]Individual, ga.PopulationSize),
		PopulationFitness: 0,
	}

	// Loop over current population by fitness
	for populationIndex := 0; populationIndex < len(population.Population); populationIndex++ {
		individual := population.GetFittest(populationIndex)

		// Create random individual to swap genes with
		randomIndividual := createIndividual(timeTable)

		// Loop over individual's genes
		for geneIndex := 0; geneIndex < len(individual.Chromosome); geneIndex++ {
			// Skip mutation if this is an elite individual
			if populationIndex > ga.ElitismCount {
				// Does this gene need mutation?
				if ga.MutationRate > rand.Float64() {
					// Swap for new gene
					individual.Chromosome[geneIndex] = randomIndividual.Chromosome[geneIndex]
				}
			}
		}

		// Add individual to population
		newPopulation.Population[populationIndex] = individual
	}

	// Return mutated population
	return newPopulation
}

/**
 * Apply crossover to population
 *
 * @param population The population to apply crossover to
 * @return The new population
 */
func (ga *GeneticAlgorithm) CrossoverPopulation(population Population) Population {
	// Create new population
	newPopulation := Population{
		Population:        make([]Individual, len(population.Population)),
		PopulationFitness: 0,
	}

	// Loop over current population by fitness
	for populationIndex := 0; populationIndex < len(population.Population); populationIndex++ {
		parent1 := population.GetFittest(populationIndex)

		// Apply crossover to this individual?
		if ga.CrossoverRate > rand.Float64() && populationIndex >= ga.ElitismCount {
			// Initialize offspring
			offspring := Individual{
				Chromosome: make([]int, len(parent1.Chromosome)),
				Fitness:    0,
			}

			// Find second parent
			parent2 := ga.SelectParent(population)

			// Loop over genome
			for geneIndex := 0; geneIndex < len(parent1.Chromosome); geneIndex++ {
				// Use half of parent1's genes and half of parent2's genes
				if 0.5 > rand.Float64() {
					offspring.Chromosome[geneIndex] = parent1.Chromosome[geneIndex]
				} else {
					offspring.Chromosome[geneIndex] = parent2.Chromosome[geneIndex]
				}
			}

			// Add offspring to new population
			newPopulation.Population[populationIndex] = offspring
		} else {
			// Add individual to new population without applying crossover
			newPopulation.Population[populationIndex] = parent1
		}
	}

	return newPopulation
}

/**
 * Calculate individual's fitness value
 *
 * @param individual
 * @param timetable
 * @return fitness
 */
func calcFitness(individual Individual, timeTable TimeTable) float64 {

	// Create new timetable object to use -- cloned from an existing timetable
	threadTimetable := timeTable

	threadTimetable.createClasses(individual)

	// Calculate fitness
	var clashes = threadTimetable.calcClashes()
	var fitness = 1 / (float64)(clashes+1)

	individual.Fitness = fitness

	return fitness
}
