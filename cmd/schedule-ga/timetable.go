package main

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

func (t *TimeTable) getNumClasses() int {
	if NumClasses > 0 {
		return NumClasses
	}

	numClasses := 0
	groups := Groups.Values.ToArray()
	for group := range groups {
		numClasses += group.GetModuleIds().Length
	}
	t.NumClasses = numClasses

	return t.NumClasses
}

func (t *TimeTable) getGroups() []Group {

	return []Group{Group{}}
}

func (t *TimeTable) getRandomTimeSlot() TimeSlot {
	timeSlotArray := getTimeSlotsValues(t)
	timeSlot := timeSlotArray[rand.Intn(949543) % len(timeSlotArray)];
	return timeSlot;
}

func (t *TimeTable) getRandomRoom() Room {
	var roomArray = getRoomsValues(t)
    var room = roomArray[rand.Intn(949543) % len(roomArray)];
    return room;
}

func (t *TimeTable) getModule(moduleID int) Module {
	return t.Modules[moduleID]
}

func getTimeSlotsValues(timeTable TimeTable) []TimeSlot{
	v := make([]TimeSlot, 0, len(timeTable.TimeSlots))
	for  _, value := range timeTable.TimeSlots {
		v = append(v, value)
	}
	return v
}

func getRoomsValues(timeTable TimeTable) []Room{
	v := make([]Room, 0, len(timeTable.Rooms))
	for  _, value := range timeTable.Rooms {
		v = append(v, value)
	}
	return v
}