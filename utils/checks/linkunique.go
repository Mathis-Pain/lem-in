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
			fmt.Printf("ERROR <linkunique.go>l13 : Wrong format. There is more than one link between the two same rooms. (%v)\n", link)
			return false
		}
	}
	return true
}
