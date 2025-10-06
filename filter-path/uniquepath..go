package filterpath

import (
	"sort"
	"strings"
)

func UniquePaths(paths [][]string) [][]string {
	seenSets := make(map[string][]string)

	for _, path := range paths {
		// extraire les salles intermédiaires sans trier
		rooms := path[1 : len(path)-1]
		key := strings.Join(rooms, ",") // clé unique pour cet ensemble

		// si on n'a pas encore enregistré ce set, on prend le chemin
		if _, exists := seenSets[key]; !exists {
			seenSets[key] = path
		} else {
			// si déjà existant, on peut garder le plus court
			if len(path) < len(seenSets[key]) {
				seenSets[key] = path
			}
		}
	}

	// récupérer les chemins filtrés
	var result [][]string
	for _, path := range seenSets {
		result = append(result, path)
	}

	// trier par longueur
	sort.Slice(result, func(i, j int) bool {
		return len(result[i]) < len(result[j])
	})

	return result
}
