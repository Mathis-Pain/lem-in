package checks

import (
	"fmt"
	"lem-in/models"
)

// Vérifie que toutes les salles présentes dans les links sont des salles existantes
func CheckLinkNames(AllRooms models.Roomlist, Links []models.Link) bool {
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
			fmt.Println("ERROR <linknames.go>-l22: Wrong format. One or more links contain an invalid room name : ", current.To)
			return false
		} else if !validName[current.From] {
			fmt.Println("ERROR <linknames.go>-l25: Wrong format. One or more links contain an invalid room name : ", current.From)
			return false
		}

	}

	return true
}
