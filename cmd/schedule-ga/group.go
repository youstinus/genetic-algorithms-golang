package main

// Group represents ...
type Group struct {
	GroupID   int
	GroupSize int
	ModuleIDs []int
}

func createGroup(groupID int, groupSize int, moduleIDs []int) (group Group) {
	group = Group{
		GroupID:   groupID,
		GroupSize: groupSize,
		ModuleIDs: moduleIDs,
	}
	return
}

func (g *Group) getModuleIDs() []int {

	return g.ModuleIDs
}

func (g *Group) getGroupID() int {

	return g.GroupID
}
