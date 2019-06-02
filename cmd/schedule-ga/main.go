package main

import (
	"fmt"
)

func main() {
	// Get a Timetable object with all the available information.
	timeTable := initializeTimetable()

	// Initialize GA
	ga := GeneticAlgorithm{
		PopulationSize: 100,
		MutationRate:   0.01,
		CrossoverRate:  0.9,
		ElitismCount:   2,
		TournamentSize: 5,
	}

	// Initialize population
	population := ga.InitPopulation(timeTable)
	//fmt.Println(population)
	// Evaluate population
	ga.EvalPopulation(&population, timeTable)
	/*fmt.Println(population)
	population.sortPopulation()
	fmt.Println(population)*/
	// Keep track of current generation
	generation := 1

	// Start evolution loop
	for !ga.IsTerminationConditionMet1(generation, 1000) && !ga.IsTerminationConditionMet2(population) {
		// Print fitness
		fmt.Println("G", generation, "Best fitness:", population.GetFittest(0).Fitness)

		// Apply crossover
		population = ga.CrossoverPopulation(population)

		// Apply mutation
		population = ga.MutatePopulation(population, timeTable)

		// Evaluate population
		ga.EvalPopulation(&population, timeTable)

		// Increment the current generation
		generation++
	}

	// Print fitness
	individ := population.GetFittest(0)
	timeTable.createClasses(&individ)
	fmt.Println()
	fmt.Println("Solution found in ", generation, " generations")
	fmt.Println("Final solution fitness: ", population.GetFittest(0).Fitness)
	fmt.Println("Clashes: ", timeTable.calcClashes())

	// Print classes
	fmt.Println()
	classes := timeTable.getClasses()
	classIndex := 1
	for _, bestClass := range classes {
		fmt.Println("Class ", classIndex, ":")
		fmt.Println("Module: ",
			timeTable.Modules[bestClass.ModuleID].Module)
		fmt.Println("Group: ",
			timeTable.Groups[bestClass.GroupID].GroupID)
		fmt.Println("Room: ",
			timeTable.Rooms[bestClass.RoomID].RoomName)
		fmt.Println("Professor: ",
			timeTable.Professors[bestClass.ProfessorID].ProfessorName)
		fmt.Println("Time: ",
			timeTable.TimeSlots[bestClass.TimeSlotID].TimeSlot)
		fmt.Println("-----")
		classIndex++
	}
}

/**
 * Creates a Timetable with all the necessary course information.
 *
 * Normally you'd get this info from a database.
 *
 * @return
 */
func initializeTimetable() TimeTable {
	// Create timeTable
	timeTable := TimeTable{
		Rooms:      make(map[int]Room),
		Professors: make(map[int]Professor),
		Modules:    make(map[int]Module),
		Groups:     make(map[int]Group),
		TimeSlots:  make(map[int]TimeSlot),
		Classes:    nil,
		NumClasses: 0,
	}

	// Set up rooms
	timeTable.AddRoom(1, "A1", 15)
	timeTable.AddRoom(2, "B1", 30)
	timeTable.AddRoom(4, "D1", 20)
	timeTable.AddRoom(5, "F1", 25)

	// Set up timeslots
	timeTable.AddTimeSlot(1, "Mon 9:00 - 11:00")
	timeTable.AddTimeSlot(2, "Mon 11:00 - 13:00")
	timeTable.AddTimeSlot(3, "Mon 13:00 - 15:00")
	timeTable.AddTimeSlot(4, "Tue 9:00 - 11:00")
	timeTable.AddTimeSlot(5, "Tue 11:00 - 13:00")
	timeTable.AddTimeSlot(6, "Tue 13:00 - 15:00")
	timeTable.AddTimeSlot(7, "Wed 9:00 - 11:00")
	timeTable.AddTimeSlot(8, "Wed 11:00 - 13:00")
	timeTable.AddTimeSlot(9, "Wed 13:00 - 15:00")
	timeTable.AddTimeSlot(10, "Thu 9:00 - 11:00")
	timeTable.AddTimeSlot(11, "Thu 11:00 - 13:00")
	timeTable.AddTimeSlot(12, "Thu 13:00 - 15:00")
	timeTable.AddTimeSlot(13, "Fri 9:00 - 11:00")
	timeTable.AddTimeSlot(14, "Fri 11:00 - 13:00")
	timeTable.AddTimeSlot(15, "Fri 13:00 - 15:00")

	// Set up professors
	timeTable.AddProfessor(1, "Dr P Smith")
	timeTable.AddProfessor(2, "Mrs E Mitchell")
	timeTable.AddProfessor(3, "Dr R Williams")
	timeTable.AddProfessor(4, "Mr A Thompson")

	// Set up modules and define the professors that teach them
	timeTable.AddModule(1, "cs1", "Computer Science", []int{1, 2})
	timeTable.AddModule(2, "en1", "English", []int{1, 3})
	timeTable.AddModule(3, "ma1", "Maths", []int{1, 2})
	timeTable.AddModule(4, "ph1", "Physics", []int{3, 4})
	timeTable.AddModule(5, "hi1", "History", []int{4})
	timeTable.AddModule(6, "dr1", "Drama", []int{1, 4})

	// Set up student groups and the modules they take.
	timeTable.AddGroup(1, 10, []int{1, 3, 4})
	timeTable.AddGroup(2, 30, []int{2, 3, 5, 6})
	timeTable.AddGroup(3, 18, []int{3, 4, 5})
	timeTable.AddGroup(4, 25, []int{1, 4})
	timeTable.AddGroup(5, 20, []int{2, 3, 5})
	timeTable.AddGroup(6, 22, []int{1, 4, 5})
	timeTable.AddGroup(7, 16, []int{1, 3})
	timeTable.AddGroup(8, 18, []int{2, 6})
	timeTable.AddGroup(9, 24, []int{1, 6})
	timeTable.AddGroup(10, 25, []int{3, 4})
	return timeTable
}
