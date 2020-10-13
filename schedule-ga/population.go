package main

import (
	"math/rand"
	"sort"
)

// Population represents ...
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

// createPopulationBySize s
// creates the initial population
func createPopulationBySize(size int) (population Population) {
	population = Population{
		Population:        make([]Individual, size),
		PopulationFitness: 0,
	}
	return
}

// createPopulationByTimeTable s
// creates the initial population
func createPopulationByTimeTable(size int, timeTable TimeTable) (population Population) {
	// Initial population
	population = Population{
		Population:        make([]Individual, size),
		PopulationFitness: 0,
	}

	// Loop over population size
	for individualCount := 0; individualCount < size; individualCount++ {
		// Create individual
		individual := createIndividual(timeTable)
		// Add individual to population
		population.Population[individualCount] = individual
	}
	return
}

// GetFittest s
/*
 * Find fittest individual in the population
 *
 * @param offset
 * @return individual Fittest individual at offset
 */
func (p *Population) GetFittest(offset int) Individual {
	// Order population by fitness
	/*Arrays.sort(_population, new Comparator<Individual>() {
		 @Override

		 public int compare(Individual o1, Individual o2)
		 {
			 if (o1.getFitness() > o2.getFitness())
			 {
				 return -1;
			 }
			 else if (o1.getFitness() < o2.getFitness())
			 {
				 return 1;
			 }
			 return 0;

		 }
	 });
	*/

	sort.Sort(ByFitness(p.Population))
	// Return the fittest individual
	//fmt.Println(p.Population[0])
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
