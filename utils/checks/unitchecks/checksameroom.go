package unitchecks

import (
	"fmt"
	"lem-in/models"
)

func CheckSameRoom(AllRooms models.Roomlist) (bool, string) {
	seenNames := make(map[string]bool)
	seenCoords := make(map[[2]int]bool)

	for _, r := range AllRooms.Rooms {
		if seenNames[r.Name] {
			return false, "duplicate name: " + r.Name
		}
		seenNames[r.Name] = true

		coord := [2]int{r.CooX, r.CooY}
		if seenCoords[coord] {
			return false, fmt.Sprintf("duplicate coordinates: (%d,%d)", r.CooX, r.CooY)
		}
		seenCoords[coord] = true
	}
	return true, ""
}
