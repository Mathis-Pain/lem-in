package models

type File struct {
	NbAnts int
	Start  string
	End    string
	Rooms  []string // List of room definitions
	Links  []string // List of link definitions

}
