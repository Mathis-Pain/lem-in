package filterpath

import "sort"

// EssentialPaths filtre les chemins pour ne garder que ceux
// qui n’utilisent pas de salles intermédiaires déjà utilisées,
// puis trie les chemins restants par longueur (court → long).
func EssentialPaths(paths [][]string) [][]string {
	var result [][]string
	used := make(map[string]bool)

	for _, path := range paths {
		ok := true
		// Vérifie les salles intermédiaires uniquement
		for _, room := range path[1 : len(path)-1] {
			if used[room] {
				ok = false
				break
			}
		}
		if ok {
			result = append(result, path)
			for _, room := range path[1 : len(path)-1] {
				used[room] = true
			}
		}
	}

	// Trier par longueur après filtrage
	sort.Slice(result, func(i, j int) bool {
		return len(result[i]) < len(result[j])
	})

	return result
}
