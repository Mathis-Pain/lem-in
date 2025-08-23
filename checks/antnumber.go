package checks

import (
	"fmt"
	"strconv"
)

// Fonction pour vérifier que le nombre de fourmis en début de fichier est valide
func CheckAntNumber(line string) bool {
	n, err := strconv.Atoi(line)

	if err != nil || n <= 0 {
		fmt.Println("ERROR : Invalid number of ants")
		return false
	}

	return true
}
