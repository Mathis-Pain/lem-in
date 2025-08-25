package unitchecks

import (
	"bufio"
	"fmt"
	"os"
)

// Fonction pour vérifier que le fichier contient bien une salle de départ (start) et une salle de fin (end)
func CheckStartEnd(content *os.File) bool {
	scanner := bufio.NewScanner(content)
	hasStart, hasEnd := false, false

	// Lit le fichier ligne par ligne pour chercher la ligne start et la ligne end
	for scanner.Scan() {
		line := scanner.Text()

		// Termine la boucle et renvoie true si les deux lignes ont été trouvée
		if hasStart && hasEnd {
			return true
		}
		if line == "##start" {
			if hasStart {
				return false
			}
			hasStart = true
			continue
		}
		if line == "##end" {
			if hasEnd {
				return false
			}
			hasEnd = true
			continue
		}
	}

	// Si la boucle se termine et qu'il manque une des deux salles, renvoie un message d'erreur
	if hasStart && !hasEnd {
		fmt.Println("ERROR : Missing an end room")
	} else if hasEnd && !hasStart {
		fmt.Println("ERROR : Missing a start room")
	}

	return false
}
