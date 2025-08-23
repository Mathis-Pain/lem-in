package checks

import (
	"bufio"
	"fmt"
	"os"
)

// Fonction générale pour vérifier que le format du fichier est correct
func FirstCheck(file *os.File) bool {
	// Vérifie qu'il y a bien une room "start" et une room "end"
	if !CheckStartEnd(file) {
		return false
	}

	// Remet le fichier à zéro pour les vérifications suivantes
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	linecount := 0

	for scanner.Scan() {
		linecount += 1
		line := scanner.Text()

		// Vérifie si la première ligne est un nombre de fourmis valide
		if linecount == 1 {
			if !CheckAntNumber(line) {
				return false
			}
			continue
		} else {
			if line[0] == 'L' {
				fmt.Println("ERROR : Invalid room name format (starting with an L)")
				return false
			}
		}
	}

	return true
}
