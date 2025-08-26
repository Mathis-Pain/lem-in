package unitchecks

import (
	"lem-in/models"
)

func LastCheck(allRooms models.Roomlist, links []models.Link) bool {

	if !Onelink(links) {
		return false
	}

	checkSameRoom, _ := CheckSameRoom(allRooms)
	if checkSameRoom {
		return true
	}
	return true
}
