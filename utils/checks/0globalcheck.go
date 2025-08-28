package checks

import (
	"bufio"
	"fmt"
	"lem-in/utils/builder"
	"os"
	"strconv"
)

// Lance toutes les fonctions de vérification du fichier
// A savoir : start et end || Nombre de fourmis valide || Nom et coordonnées des salles || Format des liens
func GlobalCheck(file *os.File) (int, bool) {
	// Vérifie qu'il y a bien une room "start" et une room "end"
	if !CheckStartEnd(file) {
		return 0, false
	}

	// Remet le fichier à zéro pour les vérifications suivantes
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	linecount, ants := 0, 0

	for scanner.Scan() {
		linecount += 1
		line := scanner.Text()

		// Vérifie si la première ligne est un nombre de fourmis valide
		if linecount == 1 {
			if !CheckAntNumber(line) {
				return ants, false
			} else {
				ants, _ = strconv.Atoi(line)
			}
		} else if line[0] == 'L' {
			// Le nom des salles ne doit pas commencer par un L (réservé aux fourmis)
<<<<<<< HEAD
			fmt.Println("ERROR <globalcheck.go>-l38: Invalid room name format (starting with an L)")
=======
			fmt.Println("ERROR <firstcheck.go>-l38: Invalid room name format (starting with an L)")
>>>>>>> refs/remotes/origin/jcolange
			return ants, false
		}
	}

	// Vérification des données des links et des salles
	return ants, CheckRoomLinkFormats(builder.ParseFile(file))
}
