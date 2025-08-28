package builder

import (
	"fmt"
	"lem-in/models"
	"os"
	"strconv"
)

func GetRoom(parts []string) models.Room {
	var current models.Room
	var err error

	current.Name = parts[0]

	current.CooX, err = strconv.Atoi(parts[1])
	RoomError(err, parts)

	current.CooY, err = strconv.Atoi(parts[2])
	RoomError(err, parts)

	return current
}

// Message d'erreur si le format n'est pas bon
func RoomError(err error, parts []string) {
	if err != nil {
		fmt.Println("ERROR <rooms.go>-l28: Invalid room coordinates for room ", parts[0])
		os.Exit(0)
	}
}
