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

	//on parcoure ces noms pour vérifier qu'il n'y aucun doublon
	for i := 0; i < len(names)-1; i++ {
		if names[i] == names[i+1] {
			fmt.Println("ERROR : Invalid format, two rooms with the same name")
			return true
		}
	}
	return true
}
