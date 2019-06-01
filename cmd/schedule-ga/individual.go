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

// creates a Individual
/*func createIndividual2(target []rune) (individual Individual) {
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
