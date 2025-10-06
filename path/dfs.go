package path

import (
	"lem-in/models"
	"sort"
)

// dfsUnique explore récursivement tous les chemins possibles sans répéter les salles intermédiaires

func dfsUnique(graph map[string][]string, current, end string, visited map[string]bool, path []string, allPaths *[][]string) {
	path = append(path, current)

	if current != end {
		visited[current] = true
	}

	if current == end {
		*allPaths = append(*allPaths, append([]string{}, path...))
	} else {
		neighbors := append([]string{}, graph[current]...) // copie des voisins
		sort.Strings(neighbors)                            // tri alphabétique pour rendre l'itération déterministe
		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				dfsUnique(graph, neighbor, end, visited, path, allPaths)
			}
		}
	}

	if current != end {
		visited[current] = false
	}
}

func FindAllPaths(graph models.Graph, start, end string) [][]string {
	var allPaths [][]string
	visited := make(map[string]bool)
	visited[start] = true

	dfsUnique(graph.Connection, start, end, visited, []string{}, &allPaths)

	sort.Slice(allPaths, func(i, j int) bool {
		return len(allPaths[i]) < len(allPaths[j])
	})
	return allPaths
}
