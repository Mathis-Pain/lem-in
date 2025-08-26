package unitchecks

import (
	"fmt"
	"lem-in/models"
)

func CheckRoomLinks(AllRooms models.Roomlist, Links []models.Link) bool {
	// Récupère toutes les salles existant dans le fichier
	validName := make(map[string]bool)

	validName[AllRooms.End.Name] = true
	validName[AllRooms.Start.Name] = true
	for _, roomName := range AllRooms.Rooms {
		validName[roomName.Name] = true
	}

	// Vérifie pour chaque Link si la salle "From" et la salle "To" existent bien dans la liste des salles
	for _, current := range Links {
		if !validName[current.To] {
			fmt.Println("ERROR : Wrong format. One or more links contain invalid room name : ", current.To)
			return false
		} else if !validName[current.From] {
			fmt.Println("ERROR : Wrong format. One or more links contain invalid room name : ", current.From)
			return false
		}

	}

	return true
}
