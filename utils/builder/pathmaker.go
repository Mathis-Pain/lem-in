package builder

import (
	"lem-in/models"
	"os"
)

// Explore récursivement les chemins
func explore(current string, end string, links []models.Link, visited map[string]bool, path []string, allPaths *[][]string) {
	// Ajouter la salle actuelle au chemin
	path = append(path, current)

	// Si on est arrivé à "end", on ajoute le chemin trouvé
	if current == end {
		// faire une copie de path car slice est réutilisé
		newRoom := make([]string, len(path))
		copy(newRoom, path)
		*allPaths = append(*allPaths, newRoom)
		return
	}

	// Marquer la salle comme visitée pour éviter les cycles
	visited[current] = true

	// Explorer tous les liens qui partent de cette salle
	for _, link := range links {
		if link.From == current && !visited[link.To] {
			explore(link.To, end, links, visited, path, allPaths)
		}
	}

	// Dé-marquer la salle (backtracking)
	visited[current] = false
}

func PathMaker(content *os.File) [][]string {
	AllRooms, Links := ParseFile(content)
	var allPaths [][]string
	visited := make(map[string]bool)

	// Lancer l’exploration depuis "start"
	explore(AllRooms.Start.Name, AllRooms.End.Name, Links, visited, []string{}, &allPaths)

	return allPaths
}
