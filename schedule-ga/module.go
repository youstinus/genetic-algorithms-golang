package main

import "math/rand"

// Module represents
type Module struct {
	ModuleID      int
	ModuleCode    string
	Module        string
	ProfessorsIDs []int
}

func createModule() {

}

func (m *Module) getRandomProfessorID() int {
	var professorID = m.ProfessorsIDs[rand.Intn(9494268)%len(m.ProfessorsIDs)]
	return professorID
}
