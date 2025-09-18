package utils

import (
	"fmt"
)

type Path struct {
	Nodes []string
}

type Ant struct {
	ID   int
	Path Path
	Step int // position actuelle dans le chemin
}

func MoveAnts() {

	// Chemins trouvés
	paths := []Path{
		{Nodes: []string{"start", "t", "E", "a", "m", "end"}},
		{Nodes: []string{"start", "h", "A", "c", "k", "end"}},
		{Nodes: []string{"start", "0", "o", "n", "e", "end"}},
	}

	nbAnts := 10
	ants := []Ant{}

	// Distribution des fourmis sur les chemins
	for i := 1; i <= nbAnts; i++ {
		path := paths[(i-1)%len(paths)]
		ants = append(ants, Ant{ID: i, Path: path, Step: 0})
	}

	// Simulation
	done := false
	for !done {
		done = true
		firstMove := true

		for i := range ants {
			if ants[i].Step < len(ants[i].Path.Nodes)-1 {
				ants[i].Step++
				room := ants[i].Path.Nodes[ants[i].Step]

				if !firstMove {
					fmt.Print(" ")
				}
				fmt.Printf("L%d-%s", ants[i].ID, room)

				firstMove = false
				done = false
			}
		}
		fmt.Println()
	}
}
