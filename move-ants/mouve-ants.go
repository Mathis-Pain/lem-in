package movement

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

// MoveAnts modifié pour écrire dans n'importe quel writer (terminal OU fichier)
func MoveAnts(paths [][]string, numAnts int, writer io.Writer) {
	// Ajoute le paramètre writer
	// Si writer est nil, utilise stdout
	if writer == nil {
		writer = os.Stdout
	}

	sortPaths(paths)
	antsPerPath := distributeAnts(paths, numAnts)

	var antIDs []int       // Identifiant de chaque fourmi
	var antPaths []int     // Le chemin choisi pour chaque fourmi
	var antPositions []int // Position actuelle de chaque fourmi dans son chemin

	antsSent := 0                         // Nombre de fourmis déjà envoyées
	antsFinished := 0                     // Nombre de fourmis arrivées à la fin
	pathCursor := make([]int, len(paths)) // Compteur de fourmis envoyées par chemin

	for antsFinished < numAnts {
		var moves []string

		// Phase 1 : Déplacer les fourmis déjà en route
		for i := 0; i < len(antIDs); i++ {
			if antPositions[i] < len(paths[antPaths[i]])-1 {
				antPositions[i]++
				room := paths[antPaths[i]][antPositions[i]]
				moves = append(moves, fmt.Sprintf("L%d-%s", antIDs[i], room))

				if antPositions[i] == len(paths[antPaths[i]])-1 {
					antsFinished++
				}
			}
		}

		// Phase 2 : Envoyer de nouvelles fourmis
		for pathIndex := 0; pathIndex < len(paths); pathIndex++ {
			if pathCursor[pathIndex] < antsPerPath[pathIndex] {
				antsSent++
				pathCursor[pathIndex]++

				antIDs = append(antIDs, antsSent)
				antPaths = append(antPaths, pathIndex)
				antPositions = append(antPositions, 0)
			}
		}

		// Écrire les mouvements (terminal OU fichier)
		if len(moves) > 0 {
			fmt.Fprintln(writer, strings.Join(moves, " "))
		}
	}
}

// on tri par longeur de chemin et si meme longueur on prend par index et i sera choisi avant j
func sortPaths(paths [][]string) {
	sort.SliceStable(paths, func(i, j int) bool {
		if len(paths[i]) == len(paths[j]) {
			return i > j
		}
		return len(paths[i]) < len(paths[j])
	})
}

// Sert à calculer à chaque fourmi affectée à un chemin la nouvelle longueur du chemin pour savoir
// si elle continue à envoyer dans ce même chemin ou si elle passe à un autre
func distributeAnts(paths [][]string, numAnts int) []int {
	antsPerPath := make([]int, len(paths))
	times := make([]int, len(paths))

	for i := range paths {
		times[i] = len(paths[i]) - 1
	}

	for ant := 0; ant < numAnts; ant++ {
		minIndex := 0
		for i := 1; i < len(times); i++ {
			if times[i] < times[minIndex] {
				minIndex = i
			}
		}

		antsPerPath[minIndex]++
		times[minIndex]++
	}

	return antsPerPath
}
