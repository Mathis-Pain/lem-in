package unitchecks

import (
	"fmt"
	"lem-in/models"
)

func Onelink(links []models.Link) bool {
	seen := make(map[models.Link]int)
	for _, link := range links {
		seen[link]++
		if seen[link] > 1 {
			fmt.Println("ERROR : Wrong format. More than one link between the two same rooms.")
			return false
		}
	}
	return true
}
