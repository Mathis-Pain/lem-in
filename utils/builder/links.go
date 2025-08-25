package builder

import (
	"fmt"
	"lem-in/models"
	"os"
	"strings"
)

func GetLink(line string) models.Link {
	var current models.Link
	parts := strings.Split(line, "-")

	if len(parts) != 2 {
		fmt.Println("ERROR : Wrong link format, must be <name1>-<name2>, is : ", current)
		os.Exit(0)
	}

	return current
}
