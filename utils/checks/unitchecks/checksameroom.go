package unitchecks

import (
	"fmt"
	"lem-in/models"
)

func CheckSameRoom(rooms models.Roomlist) bool {
	names := []string{}
	//on crée une slice de string names dans laquelle on range les noms de toutes les pièces
	for _, room := range rooms.Rooms {
		names = append(names, room.Name)
	}
	names = append(names, rooms.End.Name, rooms.Start.Name)

	seen := make(map[string]int)

	for _, name := range names {
		seen[name]++
		if seen[name] > 1 {
			fmt.Print("ERROR : Wrong format. The same room was created twice.")
			return false
		}
	}
	return true

}
