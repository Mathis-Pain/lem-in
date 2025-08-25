package unitchecks

import (
	"fmt"
	"strconv"
)

// Fonction pour vérifier que le nombre de fourmis en début de fichier est valide
func CheckAntNumber(line string) bool {
	n, err := strconv.Atoi(line)

	if err != nil || n <= 0 {
		fmt.Println("ERROR : Invalid number of ants. Must a number > 0, is", line)
		return false
	}

	return true
}
