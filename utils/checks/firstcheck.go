package checks

import (
	"bufio"
	"fmt"
	"lem-in/utils/builder"
	"lem-in/utils/checks/unitchecks"
	"os"
	"strconv"
	"strings"
)

// Fonction générale pour vérifier que le format du fichier est correct
func FirstCheck(file *os.File) (int, bool) {
	// Vérifie qu'il y a bien une room "start" et une room "end"
	if !unitchecks.CheckStartEnd(file) {
		return 0, false
	}

	// Remet le fichier à zéro pour les vérifications suivantes
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)
	linecount := 0
	n := 0

	for scanner.Scan() {
		linecount += 1
		line := scanner.Text()

		// Vérifie si la première ligne est un nombre de fourmis valide
		if linecount == 1 {
			if !unitchecks.CheckAntNumber(line) {
				return n, false
			}
			n, _ = strconv.Atoi(line)
			continue
		} else if line[0] == 'L' {
			fmt.Println("ERROR <firstcheck.go>-l38: Invalid room name format (starting with an L)")
			return n, false
			//on vérifie la validité des links
		} else if strings.Contains(line, "-") {
			if !unitchecks.CheckLinks(line) {
				fmt.Println("ERROR <firstcheck.go>-l38: Invalid link (each links needs two different rooms)")
				return n, false
			}
		}

	}

	return n, unitchecks.LastCheck(builder.ParseFile(file))
}
