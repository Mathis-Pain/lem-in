package checks

import (
	"fmt"
	"lem-in/models"
)

func CheckLinkTwoRooms(links []models.Link) bool {
	for _, link := range links {
		if link.To == link.From {
			fmt.Printf("ERROR <linktworooms.go>-l10: Invalid link format. Links must be between two different rooms. (%v)\n", link)
			return false
		}

	}

	return true
}
