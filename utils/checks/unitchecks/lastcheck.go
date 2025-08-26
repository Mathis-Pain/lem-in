package unitchecks

import (
	"lem-in/models"
)

func LastCheck(allRooms models.Roomlist, links []models.Link) bool {

	if !Onelink(links) {
		return false
	}

	if CheckSameRoom(allRooms) {
		return true
	}
	return true
}
