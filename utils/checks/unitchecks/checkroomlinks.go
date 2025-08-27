package unitchecks

import (
	"fmt"
	"lem-in/models"
)

func CheckRoomLinks(AllRooms models.Roomlist, Links []models.Link) bool {
	var names []string

	names = append(names, AllRooms.End.Name)
	names = append(names, AllRooms.Start.Name)

	for _, roomName := range AllRooms.Rooms {
		names = append(names, roomName.Name)
	}

	validName := make(map[string]bool)

	for _, v := range names {
		validName[v] = true
	}

	for _, current := range Links {
		if !validName[current.To] {
			fmt.Println("ERROR <checkroomlinks.go>-l26: Wrong format. One or more links contain invalid room name : ", current.To)
			return false
		} else if !validName[current.From] {
			fmt.Println("ERROR <checkroomlinks.go>-l29: Wrong format. One or more links contain invalid room name : ", current.From)
			return false
		}

	}

	return true
}
