package data

import (
	"fmt"
	"lem-in/models"
	"strconv"
	"strings"
)

func ExtractFile(scanExemple []string) models.File {
	var file models.File
	for i, line := range scanExemple {
		// ligne vide on continu
		if line == "" {
			continue
		}
		if line[0] == '#' && line[1] != '#' {
			continue
		}
		// on recupere le nombre de fourmi
		if i == 0 {
			NbrAnts, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println("Le nombre de fourmis n'est pas valide", err)
				return file
			} else {
				file.NbAnts = NbrAnts
			}
		}
		// on recupere la salle start
		if line == "##start" && i+1 < len(scanExemple) {
			parts := strings.Fields(scanExemple[i+1])
			file.Start = parts[0]
		}
		// on recupere" la salle end
		if line == "##end" && i+1 < len(scanExemple) {
			parts := strings.Fields(scanExemple[i+1])
			file.End = parts[0]
		}
		// on recupere les salles
		if len(strings.Fields(line)) == 3 && !strings.Contains(line, "-") {
			file.Rooms = append(file.Rooms, line)
		}
		// on recupere les liens
		if strings.Contains(line, "-") {
			file.Links = append(file.Links, line)
		}
	}
	return file
}
