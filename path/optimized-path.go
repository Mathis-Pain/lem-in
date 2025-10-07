package path

import (
	"sort"
)

const MaxInt = int(^uint(0) >> 1) // Constante pour représenter un très grand entier

// SelectPathsOptimizedWithAnts sélectionne les chemins optimaux pour minimiser le nombre de tours
func SelectPathsOptimizedWithAnts(allPaths [][]string, numAnts int) [][]string {
	if len(allPaths) == 0 {
		return nil
	}

	// Trier les chemins par longueur croissante
	sort.Slice(allPaths, func(i, j int) bool {
		return len(allPaths[i]) < len(allPaths[j])
	})

	var bestCombination [][]string
	bestTurns := MaxInt

	exploreCombinations(allPaths, [][]string{}, 0, numAnts, &bestCombination, &bestTurns)

	return bestCombination
}

// exploreCombinations explore récursivement toutes les combinaisons de chemins disjoints
func exploreCombinations(allPaths [][]string, current [][]string, startIndex int, numAnts int, best *[][]string, bestTurns *int) {
	if len(current) > 0 {
		turns := calculateTurns(current, numAnts)
		if turns < *bestTurns {
			*bestTurns = turns
			*best = copyPaths(current)
		}
	}

	for i := startIndex; i < len(allPaths); i++ {
		newPath := allPaths[i]
		if isDisjointWithAll(current, newPath) {
			newCurrent := copyPaths(current)
			newCurrent = append(newCurrent, newPath)
			exploreCombinations(allPaths, newCurrent, i+1, numAnts, best, bestTurns)
		}
	}
}

// isDisjointWithAll vérifie si un chemin est disjoint avec tous les chemins existants
func isDisjointWithAll(paths [][]string, newPath []string) bool {
	// Créer un set de toutes les salles intermédiaires déjà présentes
	occupied := make(map[string]bool)
	for _, path := range paths {
		for i := 1; i < len(path)-1; i++ { // exclure start/end
			occupied[path[i]] = true
		}
	}

	// Vérifier si newPath partage une salle
	for i := 1; i < len(newPath)-1; i++ {
		if occupied[newPath[i]] {
			return false
		}
	}
	return true
}

// copyPaths fait une copie profonde d'une slice de chemins
func copyPaths(paths [][]string) [][]string {
	cp := make([][]string, len(paths))
	for i := range paths {
		cp[i] = append([]string{}, paths[i]...)
	}
	return cp
}

// calculateTurns calcule le nombre de tours nécessaires pour faire passer toutes les fourmis
func calculateTurns(paths [][]string, numAnts int) int {
	if len(paths) == 0 {
		return MaxInt
	}

	pathLengths := make([]int, len(paths))
	for i, path := range paths {
		pathLengths[i] = len(path) - 1
	}

	antCounts := make([]int, len(paths))
	finishTimes := make([]int, len(paths))
	copy(finishTimes, pathLengths)

	for ant := 0; ant < numAnts; ant++ {
		// Chemin avec le temps de fin le plus court
		minIndex := 0
		for i := 1; i < len(finishTimes); i++ {
			if finishTimes[i] < finishTimes[minIndex] {
				minIndex = i
			}
		}

		antCounts[minIndex]++
		finishTimes[minIndex] = pathLengths[minIndex] + antCounts[minIndex] - 1
	}

	maxTime := 0
	for _, t := range finishTimes {
		if t > maxTime {
			maxTime = t
		}
	}

	return maxTime
}
