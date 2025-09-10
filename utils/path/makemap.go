package path

import "lem-in/models"

// Transforme la liste des salles en carte plus lisible et plus performante
func MakeMap(allRooms *models.Roomlist) (map[string]*models.Room, string, string) {
	// Crée une map qui associe chaque nom de salle à sa structure pour pouvoir rapidement retrouver les salles sans parcourir des struct
	roomsMap := make(map[string]*models.Room)

	// Ajoute toutes les salles à la map
	roomsMap[allRooms.Start.Name] = &allRooms.Start
	roomsMap[allRooms.End.Name] = &allRooms.End
	for i := range allRooms.Rooms {
		roomsMap[allRooms.Rooms[i].Name] = &allRooms.Rooms[i]
	}

	startRoom := allRooms.Start.Name
	endRoom := allRooms.End.Name

	return roomsMap, startRoom, endRoom
}
