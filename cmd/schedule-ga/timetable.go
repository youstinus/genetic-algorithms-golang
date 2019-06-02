package main

import (
	"fmt"
	"math/rand"
)

// TimeTable repre
type TimeTable struct {
	Rooms      map[int]Room
	Professors map[int]Professor
	Modules    map[int]Module
	Groups     map[int]Group
	TimeSlots  map[int]TimeSlot
	Classes    []Class
	NumClasses int
}

func createTimeTable() {

}

/**
 * Add new room
 *
 * @param roomId
 * @param roomName
 * @param capacity
 */
func (t *TimeTable) AddRoom(roomID int, roomName string, capacity int) {
	t.Rooms[roomID] = Room{
		RoomID:   roomID,
		RoomName: roomName,
		Capacity: capacity,
	}
}

/**
 * Add new timeslot
 *
 * @param timeslotId
 * @param timeslot
 */
func (t *TimeTable) AddTimeSlot(timeSlotID int, timeSlot string) {
	t.TimeSlots[timeSlotID] = TimeSlot{
		TimeSlotID: timeSlotID,
		TimeSlot:   timeSlot,
	}
}

/**
 * Add new professor
 *
 * @param professorId
 * @param professorName
 */
func (t *TimeTable) AddProfessor(professorID int, professorName string) {
	t.Professors[professorID] = Professor{
		ProfessorID:   professorID,
		ProfessorName: professorName,
	}
}

/**
 * Add new module
 *
 * @param moduleId
 * @param moduleCode
 * @param module
 * @param professorIds
 */
func (t *TimeTable) AddModule(moduleID int, moduleCode string, module string, professorIDs []int) {
	t.Modules[moduleID] = Module{
		ModuleID:      moduleID,
		ModuleCode:    moduleCode,
		Module:        module,
		ProfessorsIDs: professorIDs,
	}
}

/**
 * Add new group
 *
 * @param groupId
 * @param groupSize
 * @param moduleIds
 */
func (t *TimeTable) AddGroup(groupID int, groupSize int, moduleIDs []int) {
	t.Groups[groupID] = Group{
		GroupID:   groupID,
		GroupSize: groupSize,
		ModuleIDs: moduleIDs,
	}
	t.NumClasses = 0
}

func (t *TimeTable) getNumClasses() int {
	if t.NumClasses > 0 {
		return t.NumClasses
	}

	numClasses := 0

	groups := t.getGroups()
	for _, group := range groups {
		numClasses += len(group.getModuleIDs())
	}
	t.NumClasses = numClasses

	return t.NumClasses
}

func (t *TimeTable) getGroups() []Group {
	keys := make([]int, 0, len(t.Groups))
	values := make([]Group, 0, len(t.Groups))

	for k, v := range t.Groups {
		keys = append(keys, k)
		values = append(values, v)
	}
	return values
}

func (t *TimeTable) getClasses() []Class {
	keys := make([]int, 0, len(t.Classes))
	values := make([]Class, 0, len(t.Classes))

	for k, v := range t.Classes {
		keys = append(keys, k)
		values = append(values, v)
	}
	return values
}

func (t *TimeTable) getRandomTimeSlot() TimeSlot {
	timeSlotArray := t.getTimeSlotsValues()
	timeSlot := timeSlotArray[rand.Intn(949543)%len(timeSlotArray)]
	return timeSlot
}

func (t *TimeTable) getRandomRoom() Room {
	var roomArray = t.getRoomsValues()
	var room = roomArray[rand.Intn(949543)%len(roomArray)]
	return room
}

func (t *TimeTable) getModule(moduleID int) Module {
	return t.Modules[moduleID]
}

func (t *TimeTable) getTimeSlotsValues() []TimeSlot {
	v := make([]TimeSlot, 0, len(t.TimeSlots))
	for _, value := range t.TimeSlots {
		v = append(v, value)
	}
	return v
}

func (t *TimeTable) getRoomsValues() []Room {
	v := make([]Room, 0, len(t.Rooms))
	for _, value := range t.Rooms {
		v = append(v, value)
	}
	return v
}

/**
 * Create classes using individual's chromosome
 *
 * One of the two important methods in this class; given a chromosome,
 * unpack it and turn it into an array of Class (with a capital C) objects.
 * These Class objects will later be evaluated by the calcClashes method,
 * which will loop through the Classes and calculate the number of
 * conflicting timeslots, rooms, professors, etc.
 *
 * While this method is important, it's not really difficult or confusing.
 * Just loop through the chromosome and create Class objects and store them.
 *
 * @param individual
 */
func (t *TimeTable) createClasses(individual *Individual) {
	// Init classes
	t.Classes = make([]Class, t.getNumClasses())

	// Get individual's chromosome
	chromosome := individual.Chromosome
	chromosomePos := 0
	classIndex := 0

	for _, group := range t.getGroups() {
		moduleIDs := group.getModuleIDs()
		for _, moduleID := range moduleIDs {
			t.Classes[classIndex] = Class{
				ClassID:  classIndex,
				GroupID:  group.getGroupID(),
				ModuleID: moduleID,
			}

			// Add timeslot
			t.Classes[classIndex].TimeSlotID = chromosome[chromosomePos]
			chromosomePos++

			// Add room
			t.Classes[classIndex].RoomID = chromosome[chromosomePos]
			chromosomePos++

			// Add professor
			t.Classes[classIndex].ProfessorID = chromosome[chromosomePos]
			chromosomePos++

			classIndex++
		}
	}
}

/**
 * Calculate the number of clashes between Classes generated by a
 * chromosome.
 *
 * The most important method in this class; look at a candidate timetable
 * and figure out how many constraints are violated.
 *
 * Running this method requires that createClasses has been run first (in
 * order to populate this.classes). The return value of this method is
 * simply the number of constraint violations (conflicting professors,
 * timeslots, or rooms), and that return value is used by the
 * GeneticAlgorithm.calcFitness method.
 *
 * There's nothing too difficult here either -- loop through this.classes,
 * and check constraints against the rest of the this.classes.
 *
 * The two inner `for` loops can be combined here as an optimization, but
 * kept separate for clarity. For small values of this.classes.length it
 * doesn't make a difference, but for larger values it certainly does.
 *
 * @return numClashes
 */
func (t *TimeTable) calcClashes() int {
	clashes := 0

	for _, classA := range t.Classes {
		// Check room capacity
		roomCapacity := t.getRoom(classA.RoomID).Capacity
		groupSize := t.getGroup(classA.GroupID).GroupSize

		if roomCapacity < groupSize {
			clashes++
		}

		// Check if room is taken
		for _, classB := range t.Classes {
			if classA.RoomID == classB.RoomID && classA.TimeSlotID == classB.TimeSlotID && classA.ClassID != classB.ClassID {
				clashes++
				break
			}
		}

		// Check if professor is available
		for _, classB := range t.Classes {
			if classA.ProfessorID == classB.ProfessorID && classA.TimeSlotID == classB.TimeSlotID && classA.ClassID != classB.ClassID {
				clashes++
				break
			}
		}
	}

	return clashes
}

/**
 * Get room from roomId
 *
 * @param roomId
 * @return room
 */
func (t *TimeTable) getRoom(roomID int) Room {
	if t.Rooms[roomID] == (Room{}) {
		fmt.Println("Rooms doesn't contain key", roomID)
	}
	return t.Rooms[roomID]
}

/**
 * Get group from groupId
 *
 * @param groupId
 * @return group
 */
func (t *TimeTable) getGroup(groupID int) Group {
	return t.Groups[groupID]
}
