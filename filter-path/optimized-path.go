package filterpath

import (
	"sort"
)

// SelectPathsOptimizedWithAnts sélectionne les chemins optimaux en fonction du nombre de fourmis
func SelectPathsOptimizedWithAnts(allPaths [][]string, numAnts int) [][]string {
	if len(allPaths) == 0 {
		return nil
	}

	// Trier les chemins par longueur
	sort.Slice(allPaths, func(i, j int) bool {
		return len(allPaths[i]) < len(allPaths[j])
	})

	// Variable pour stocker la meilleure combinaison
	var bestCombination [][]string
	bestTurns := int(^uint(0) >> 1) // Max int

	// Explorer toutes les combinaisons possibles de chemins disjoints
	exploreCombinations(allPaths, [][]string{}, 0, numAnts, &bestCombination, &bestTurns)

	return bestCombination
}

// exploreCombinations explore récursivement toutes les combinaisons de chemins disjoints
func exploreCombinations(allPaths [][]string, current [][]string, startIdx int, numAnts int, best *[][]string, bestTurns *int) {
	// Évaluer la combinaison actuelle si elle contient au moins un chemin
	if len(current) > 0 {
		turns := calculateTurns(current, numAnts)
		if turns < *bestTurns {
			*bestTurns = turns
			*best = make([][]string, len(current))
			copy(*best, current)
		}
	}

	// Essayer d'ajouter chaque chemin restant
	for i := startIdx; i < len(allPaths); i++ {
		newPath := allPaths[i]

		// Vérifier si ce chemin est disjoint avec tous les chemins déjà sélectionnés
		if isDisjointWithAll(current, newPath) {
			// Ajouter ce chemin et continuer l'exploration
			newCurrent := append(append([][]string{}, current...), newPath)
			exploreCombinations(allPaths, newCurrent, i+1, numAnts, best, bestTurns)
		}
	}
}

// isDisjointWithAll vérifie si un chemin est disjoint avec tous les chemins de la liste
func isDisjointWithAll(paths [][]string, newPath []string) bool {
	for _, path := range paths {
		if !areDisjoint(path, newPath) {
			return false
		}
	}
	return true
}

// areDisjoint vérifie si deux chemins sont disjoints (ne partagent pas de salles intermédiaires)
func areDisjoint(path1, path2 []string) bool {
	// Créer un set des salles intermédiaires du premier chemin (sans start et end)
	intermediateRooms := make(map[string]bool)
	for i := 1; i < len(path1)-1; i++ {
		intermediateRooms[path1[i]] = true
	}

	// Vérifier si le second chemin partage des salles intermédiaires
	for i := 1; i < len(path2)-1; i++ {
		if intermediateRooms[path2[i]] {
			return false
		}
	}

	return true
}

// calculateTurns calcule le nombre de tours nécessaires pour faire passer toutes les fourmis
func calculateTurns(paths [][]string, numAnts int) int {
	if len(paths) == 0 {
		return int(^uint(0) >> 1) // Max int
	}

	// Calculer les longueurs de chaque chemin (nombre de mouvements)
	pathLengths := make([]int, len(paths))
	for i, path := range paths {
		pathLengths[i] = len(path) - 1 // -1 car on compte les mouvements, pas les salles
	}

	// Distribution optimale des fourmis
	// On utilise une approche greedy : on assigne chaque fourmi au chemin qui finira le plus tôt
	antCounts := make([]int, len(paths))   // Nombre de fourmis par chemin
	finishTimes := make([]int, len(paths)) // Temps de fin pour chaque chemin

	// Initialiser les temps de fin avec la longueur du premier ant sur chaque chemin
	for i := range finishTimes {
		finishTimes[i] = pathLengths[i]
	}
	antCounts[0] = 1 // La première fourmi prend le chemin le plus court

	// Distribuer les fourmis restantes
	for ant := 1; ant < numAnts; ant++ {
		// Trouver le chemin avec le temps de fin le plus court
		minIdx := 0
		minTime := finishTimes[0]
		for i := 1; i < len(finishTimes); i++ {
			if finishTimes[i] < minTime {
				minTime = finishTimes[i]
				minIdx = i
			}
		}

		// Assigner cette fourmi au chemin trouvé
		antCounts[minIdx]++
		// Le nouveau temps de fin = longueur du chemin + nombre de fourmis sur ce chemin - 1
		finishTimes[minIdx] = pathLengths[minIdx] + antCounts[minIdx] - 1
	}

	// Le temps total est le maximum des temps de fin
	maxTime := 0
	for _, time := range finishTimes {
		if time > maxTime {
			maxTime = time
		}
	}

	return maxTime
}
