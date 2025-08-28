package checks

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

		switch line {
		case "##start":
			if hasStart {
				fmt.Println("ERROR <startend.go>-l21: There are multiple start rooms.")
				return false
			} else {
				hasStart = true
			}
		case "##end":
			if hasEnd {
				fmt.Println("ERROR <startend.go>-l28: There are multiple end rooms.")
				return false
			} else {
				hasEnd = true
			}
		}
	}

	// Une fois le fichier lu en entier, vérifie que les deux salles sont bien présentes
	if hasStart && hasEnd {
		return true
	}

	// Si la boucle se termine et qu'il manque une des deux salles, renvoie un message d'erreur
	if hasStart && !hasEnd {
		fmt.Println("ERROR <startend.go>-l40: Missing an end room.")
	} else if hasEnd && !hasStart {
		fmt.Println("ERROR <startend.go>-l42: Missing a start room.")
	}

	return false
}
