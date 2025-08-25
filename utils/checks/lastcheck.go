package checks

import (
	"lem-in/models"
	"lem-in/utils/checks/unitchecks"
)

func LastCheck(allRooms models.Roomlist, links []models.Link) bool {

	if !unitchecks.Onelink(links) {
		return false
	}

	if unitchecks.CheckSameRoom(allRooms) {
		return false
	}
	return true
}
