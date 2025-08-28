package checks

import (
	"fmt"
	"lem-in/models"
)

func CheckLinkUnique(links []models.Link) bool {
	seen := make(map[models.Link]int)
	for _, link := range links {
		seen[link]++
		if seen[link] > 1 {
			fmt.Println("ERROR <linkunique.go>l13 : Wrong format. More than one link between the two same rooms.")
			return false
		}
	}
	return true
}
