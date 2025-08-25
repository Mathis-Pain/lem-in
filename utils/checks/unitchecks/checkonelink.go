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
			fmt.Println("Error: there is more than one way")
			return false
		}
	}
	return true
}
