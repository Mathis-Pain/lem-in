package checks

import (
	"lem-in/models"
)

func CheckRoomLinkFormats(allRooms models.Roomlist, links []models.Link) bool {

	if !CheckLinkUnique(links) || !CheckRoomUnique(allRooms) || !CheckLinkNames(allRooms, links) || !CheckLinkTwoRooms(links) {
		return false
	}

	return true
}
