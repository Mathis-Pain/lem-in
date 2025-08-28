package checks

import (
	"fmt"
	"lem-in/models"
)

func CheckLinkTwoRooms(links []models.Link) bool {
	for _, link := range links {
		if link.To == link.From {
			fmt.Println("ERROR <linktworooms.go>-l10: Invalid link (each links needs two different rooms)")
			return false
		}

	}

	return true
}
