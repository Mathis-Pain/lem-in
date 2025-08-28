package move

import (
	"fmt"
	"lem-in/models"
)

func OneStep(path []string, Ants []models.Ant, finished int) ([]models.Ant, int) {
	end := path[len(path)-1]
	var movements string
	for j := 0; j < len(Ants); j++ {
		for i := 0; i < len(path); i++ {
			switch Ants[j].Position {
			case end:
				finished += 1
				goto nextAnt
			case path[i]:
				Ants[j].Position = path[i+1]
				movements += fmt.Sprint(Ants[j].Name, " ", Ants[j].Position, " ")
				goto nextAnt
			}
		}
	nextAnt:
	}
	fmt.Println(movements)
	return Ants, finished
}

func CreateAnt(allPaths models.PathList, i int) models.Ant {
	var currentAnt models.Ant
	currentAnt.Name = fmt.Sprintf("L%v", i)
	currentAnt.Position = allPaths[0][0]

	return currentAnt
}

func MoveAnts(allPaths models.PathList, allAnts int) {
	var Ants []models.Ant
	finished := 0

	for i := 1; i <= allAnts; i++ {
		Ants = append(Ants, CreateAnt(allPaths, i))
	}

	for finished <= allAnts {

		for i := 1; i < len(allPaths); i++ {
			Ants, finished = OneStep(allPaths[i], Ants, finished)
		}
	}

}
