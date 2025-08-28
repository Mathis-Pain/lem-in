package checks

import (
	"fmt"
	"strconv"
)

// Fonction pour vérifier que le nombre de fourmis en début de fichier est valide
func CheckAntNumber(line string) bool {
	ants, err := strconv.Atoi(line)

	if err != nil || ants <= 0 {
		fmt.Println("ERROR <antnumber.go>-l13: Invalid number of ants. Must be a number > 0, is", line)
		return false
	}

	return true
}
