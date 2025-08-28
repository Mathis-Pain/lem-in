package checks

import (
	"fmt"
	"lem-in/models"
)

func CheckRoomUnique(AllRooms models.Roomlist) bool {
	seenNames := make(map[string]bool)
	seenCoords := make(map[[2]int]bool)

	for _, r := range AllRooms.Rooms {
		if seenNames[r.Name] {
			fmt.Printf("ERROR <roomunique.go>-l14: Wrong format. Two rooms have the same name : %v\n", r.Name)
			return false
		}
		seenNames[r.Name] = true

		coord := [2]int{r.CooX, r.CooY}
		if seenCoords[coord] {
			fmt.Printf("ERROR <roomunique.go>-l21: Wrong format. Two rooms have the same coordinates: %d,%d\n", r.CooX, r.CooY)
			return false
		}
		seenCoords[coord] = true
	}
	return true
}
