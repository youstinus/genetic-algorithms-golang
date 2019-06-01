package main

// Class that represents ...
type Class struct {
	ClassID     int
	GroupID     int
	ModuleID    int
	ProfessorID int
	TimeSlotID  int
	RoomID      int
}

func createClass(classID int, groupID int, moduleID int) (class Class) {

	class = Class{
		ClassID:  classID,
		GroupID:  groupID,
		ModuleID: moduleID,
	}
	return
}
