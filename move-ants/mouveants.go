package movement

import (
	"fmt"
	"strings"
)

// SimulateMovement affiche le déplacement des fourmis tour par tour
func MoveAnts(paths [][]string, numAnts int) {
	// Trier les chemins : par longueur croissante, et inverser si même longueur
	// (pour avoir l'ordre t, h, 0 au lieu de 0, h, t)
	sortPaths(paths)

	// Distribuer les fourmis sur les chemins
	antsPerPath := distributeAnts(paths, numAnts)

	// Listes pour suivre les fourmis actives
	var antIDs []int       // ID de chaque fourmi
	var antPaths []int     // Chemin assigné à chaque fourmi
	var antPositions []int // Position actuelle dans le chemin (0 = start)

	antsSent := 0     // Nombre de fourmis envoyées
	antsFinished := 0 // Nombre de fourmis arrivées

	// Compteur de fourmis envoyées par chemin
	pathCursor := make([]int, len(paths))

	// Boucle tant que toutes les fourmis ne sont pas arrivées
	for antsFinished < numAnts {
		var moves []string

		// Phase 1 : Déplacer les fourmis déjà en route
		for i := 0; i < len(antIDs); i++ {
			// Si la fourmi n'est pas encore arrivée
			if antPositions[i] < len(paths[antPaths[i]])-1 {
				// Avancer d'une position
				antPositions[i]++
				room := paths[antPaths[i]][antPositions[i]]
				moves = append(moves, fmt.Sprintf("L%d-%s", antIDs[i], room))

				// Vérifier si elle est arrivée
				if antPositions[i] == len(paths[antPaths[i]])-1 {
					antsFinished++
				}
			}
		}

		// Phase 2 : Envoyer de nouvelles fourmis
		for pathIdx := range paths {
			// Si on doit encore envoyer des fourmis sur ce chemin
			if pathCursor[pathIdx] < antsPerPath[pathIdx] {
				antsSent++
				pathCursor[pathIdx]++

				// Créer une nouvelle fourmi
				antIDs = append(antIDs, antsSent)
				antPaths = append(antPaths, pathIdx)
				antPositions = append(antPositions, 0) // Position 0 = start
			}
		}

		// Afficher les mouvements
		if len(moves) > 0 {
			fmt.Println(strings.Join(moves, " "))
		}
	}
}

// distributeAnts distribue les fourmis sur les chemins et retourne le nombre de fourmis par chemin
func distributeAnts(paths [][]string, numAnts int) []int {
	antsPerPath := make([]int, len(paths))
	times := make([]int, len(paths))

	// Initialiser avec les longueurs des chemins
	for i := range paths {
		times[i] = len(paths[i]) - 1
	}

	// Assigner chaque fourmi au chemin qui finit le plus tôt
	for ant := 0; ant < numAnts; ant++ {
		// Trouver le chemin avec le temps minimum
		minIdx := 0
		for i := 1; i < len(times); i++ {
			if times[i] < times[minIdx] {
				minIdx = i
			}
		}

		antsPerPath[minIdx]++
		times[minIdx]++
	}

	return antsPerPath
}

// sortPaths trie les chemins par longueur, et inverse l'ordre pour les chemins de même longueur
func sortPaths(paths [][]string) {
	// Créer une liste avec index pour garder trace de l'ordre original
	type pathWithIndex struct {
		path    []string
		origIdx int
		length  int
	}

	items := make([]pathWithIndex, len(paths))
	for i := range paths {
		items[i] = pathWithIndex{
			path:    paths[i],
			origIdx: i,
			length:  len(paths[i]),
		}
	}

	// Trier par longueur croissante, puis par index décroissant si même longueur
	for i := 0; i < len(items)-1; i++ {
		for j := i + 1; j < len(items); j++ {
			shouldSwap := false

			if items[i].length > items[j].length {
				// i est plus long que j, on swap
				shouldSwap = true
			} else if items[i].length == items[j].length && items[i].origIdx < items[j].origIdx {
				// Même longueur, on inverse l'ordre
				shouldSwap = true
			}

			if shouldSwap {
				items[i], items[j] = items[j], items[i]
			}
		}
	}

	// Réappliquer dans paths
	for i := range paths {
		paths[i] = items[i].path
	}
}
