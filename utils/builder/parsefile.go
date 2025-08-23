package builder

import (
	"bufio"
	"fmt"
	"lem-in/models"
	"lem-in/utils/checks"
	"os"
	"strings"
)

func ParseFile(file *os.File) (models.Roomlist, []models.Link) {
	file.Seek(0, 0)

	var AllRooms models.Roomlist
	var Links []models.Link
	var NextIsStart, NextIsEnd bool
	var NoMoreRooms bool

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		// Vérifie si la prochaine ligne est la salle de départ ou la salle de fin.
		if line == "##start" {
			NextIsStart = true
			continue
		} else if line == "##end" {
			NextIsEnd = true
			continue
		} else if line[0] == '#' {
			// Si la ligne n'est ni start ni end mais commence par un #, elle est considérée comme un commentaire et ignorée
			continue
		}

		// Découpe la ligne au niveau des espaces
		parts := strings.Split(line, " ")

		// Vérifie si une pièce ne se trouve pas au mauvais endroit du fichier
		// Il ne peut pas y avoir de nouvelle pièce une fois qu'on a commencé à ajouter des links
		if NoMoreRooms && len(parts) == 3 {
			fmt.Println("ERROR : Wrong file format : new room added after the end of the room list")
			os.Exit(0)
		}

		// Si la ligne comporte bien trois parties (noms, x et y), ajoute la salle à la liste des salles existantes
		// La salle de départ et la salle de fin sont ajoutées dans des variables à part
		if len(parts) == 3 {
			if NextIsStart {
				AllRooms.Start = GetRoom(parts)
				NextIsStart = false
			} else if NextIsEnd {
				AllRooms.End = GetRoom(parts)
				NextIsEnd = false
			} else {
				AllRooms.Rooms = append(AllRooms.Rooms, GetRoom(parts))
			}
			continue
		} else {
			NoMoreRooms = true
			// Sinon, ça veut dire qu'on a atteint la fin de la liste de salles
		}

		// Si la ligne ne comporte pas trois parties, on vérifie s'il s'agit d'un "link", tunnel entre deux salles
		if strings.Contains(line, "-") && checks.CheckLinks(line) {
			Links = append(Links, GetLink(line))

		}
	}

	return AllRooms, Links
}
