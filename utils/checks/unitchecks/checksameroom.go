package unitchecks

import (
	"fmt"
	"lem-in/models"
)

func CheckSameRoom(AllRooms models.Roomlist) bool {
	seenNames := make(map[string]bool)
	seenCoords := make(map[[2]int]bool)

	for _, r := range AllRooms.Rooms {
		if seenNames[r.Name] {
			fmt.Printf("ERROR <checksameroom.go>-l14: Wrong format. Same room name : %v", r.Name)
			return false
		}
		seenNames[r.Name] = true

		coord := [2]int{r.CooX, r.CooY}
		if seenCoords[coord] {
			fmt.Printf("ERROR <checksameroom.go>-l21: Wrong format. Same room coordinates: %d,%d", r.CooX, r.CooY)
			return false
		}
		seenCoords[coord] = true
	}
	return true
}
