package data

import (
	"fmt"
	"lem-in/models"
	"strconv"
	"strings"
)

func TestFile(file models.File) (bool, models.File) {
	// on verifie qu'il n'y ai qu'un start
	if len(file.Start) == 0 {
		fmt.Println("Erreur format fichier start")
		return false, file
	}
	// on verifie qu'il n'y au qu'un end
	if len(file.End) == 0 {
		fmt.Println("Erreur format de fichier end")
		return false, file
	}
	// on verifie que le nombre de fourmis soit valide
	if file.NbAnts <= 0 || file.NbAnts > 1001 {
		fmt.Println("Erreur nombre de formis invalide")
		return false, file
	}
	// on verifie les salles
	seenNameRoom := make(map[string]int)
	seenCoordRoom := make(map[string]int)
	for _, room := range file.Rooms {
		// on verifie que les coordonnées de salle soit valide
		testroom := strings.Fields(room)
		_, err1 := strconv.Atoi(testroom[1])
		_, err2 := strconv.Atoi(testroom[2])
		if err1 != nil || err2 != nil {
			fmt.Println("Erreur coordonnées de room invalide")
			return false, file
		}
		seenNameRoom[testroom[0]]++
		testroomkey := testroom[1] + "," + testroom[2]
		seenCoordRoom[testroomkey]++
		//on verifie que le nom de la salle soit present une seule fois
		if seenNameRoom[testroom[0]] > 1 {
			fmt.Println("Erreur Nom de salle deja utilisé")
			return false, file
		}
		// on verifie que les coordonnée soit presente une seule fois
		if seenCoordRoom[testroomkey] > 1 {
			fmt.Println("=Erreur les coordonées sont déja utilisées")
			return false, file
		}
	}
	// on verifie les liens
	presentRoomG := false
	presentRoomD := false
	// on verifie que les salles des liens existe dans la liste des salles
	for _, link := range file.Links {

		salles := strings.Split(link, "-")
		for _, room := range file.Rooms {
			roomName := strings.Fields(room)
			if salles[0] == roomName[0] {
				presentRoomG = true
			}
			if salles[1] == roomName[0] {
				presentRoomD = true
			}
			if presentRoomG && presentRoomD {
				break

			}
		}
		if !presentRoomG || !presentRoomD {

			fmt.Println("Erreur une salle n'est pas presente dans la liste des salle dispinible")
			return false, file
		}
	}
	return true, file
}

// si on arrive ici les formats sont valide
