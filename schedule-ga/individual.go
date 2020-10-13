package main

// Individual for this genetic algorithm
type Individual struct {
	Chromosome []int
	Fitness    float64
}

func createIndividual(timeTable TimeTable) (individual Individual) {
	numClasses := timeTable.getNumClasses()
	chromosomeLength := numClasses * 3
	newChromosome := make([]int, chromosomeLength)
	chromosomeIndex := 0

	for _, group := range timeTable.getGroups() {

		for _, moduleID := range group.getModuleIDs() {
			timeSlotID := timeTable.getRandomTimeSlot().TimeSlotID
			newChromosome[chromosomeIndex] = timeSlotID
			chromosomeIndex++

			// Add random room
			roomID := timeTable.getRandomRoom().RoomID
			newChromosome[chromosomeIndex] = roomID
			chromosomeIndex++

			// Add random professor
			module := timeTable.getModule(moduleID)
			newChromosome[chromosomeIndex] = module.getRandomProfessorID()
			chromosomeIndex++
		}
	}

	individual = Individual{
		Chromosome: newChromosome,
		Fitness:    0,
	}
	return
}

// crosses over 2 Individuals
/*func crossover(d1 Individual, d2 Individual) Individual {
	child := Individual{
		Chromosome: make([]int, len(d1.Chromosome)),
		Fitness:    0,
	}
	mid := rand.Intn(len(d1.Chromosome))
	for i := 0; i < len(d1.Chromosome); i++ {
		if i > mid {
			child.Chromosome[i] = d1.Chromosome[i]
		} else {
			child.Chromosome[i] = d2.Chromosome[i]
		}

	}
	return child
}

// mutate the Individual
func (i *Individual) mutate() {
	for d := 0; d < len(i.Chromosome); d++ {
		if rand.Float64() < MutationRate {
			i.Chromosome[d] = int(Letters[rand.Intn(33)])
		}
	}
}

// calcFitness calculates fitness
func (i *Individual) calcFitness(target) {

}

// creates a Individual
func createIndividual2(target []rune) (individual Individual) {
	ba := make([]rune, len(target))
	for i := 0; i < len(target); i++ {
		ba[i] = rune(Letters[rand.Intn(33)])
	}
	individual = Individual{
		Chromosome: ba,
		Fitness:    0,
	}
	individual.calcFitness(target)
	return
}*/

// calculates the fitness of the Individual
/*func (d *Individual) calcFitness(target []rune) {
	score := 0
	for i := 0; i < len(d.Chromosome); i++ {
		if d.Chromosome[i] == target[i] {
			score++
		}
	}
	d.Fitness = float64(score) / float64(len(d.Chromosome))
	return
}*/
