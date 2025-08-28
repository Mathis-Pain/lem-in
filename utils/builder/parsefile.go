package builder

import (
	"bufio"
	"fmt"
	"lem-in/models"
	"os"
	"strings"
)

func ParseFile(file *os.File) (models.Roomlist, []models.Link) {
	file.Seek(0, 0)
	var AllRooms models.Roomlist
	var Links []models.Link
	var NextIsStart, NextIsEnd bool
	NoMoreRooms := false

	scanner := bufio.NewScanner(file)
	linecount := 0

	for scanner.Scan() {
		line := scanner.Text()
		linecount++
		// Vérifie si la prochaine ligne est la salle de départ ou la salle de fin.
		if line == "##start" {
			NextIsStart = true
			continue
		} else if line == "##end" {
			NextIsEnd = true
			continue
		} else if line[0] == '#' || linecount == 1 {
			// Si la ligne n'est ni start ni end mais commence par un #, elle est considérée comme un commentaire et ignorée
			// Passe la première ligne
			continue
		}

		// Découpe la ligne au niveau des espaces
		parts := strings.Split(line, " ")

		// Vérifie si une pièce ne se trouve pas au mauvais endroit du fichier
		// Il ne peut pas y avoir de nouvelle pièce une fois qu'on a commencé à ajouter des links
		if NoMoreRooms && len(parts) == 3 {
			fmt.Printf("ERROR <parsefile.go>-l43 : Wrong file format : new room added after the end of the room list (line number %d)\n", linecount)
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
		} else if strings.Contains(line, "-") {
			// Si la ligne ne comporte pas trois parties, on vérifie s'il s'agit d'un "link", tunnel entre deux salles
			NoMoreRooms = true
			Links = append(Links, GetLink(line))
		} else {
			fmt.Printf("ERROR <parsefile.go>-l63 : Wrong file format : line is not a comment, a room or a link. (%v - line number %v)\n", line, linecount)
			os.Exit(1)
		}
	}

	return AllRooms, Links
}
