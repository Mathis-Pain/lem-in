package path

import (
	"container/list"
	"lem-in/models"
)

// Fonction pour implémenter l'algorithme Edmonds-Karp et mettre à jour la liste des chemins
func EdmondsKarp(allRooms *models.Roomlist, links []models.Link) [][]string {
	roomsMap, startRoom, endRoom := MakeMap(allRooms)
	capacity, flow := BuildGraph(roomsMap, startRoom, endRoom, links)

	var paths [][]string
	for {
		parent := make(map[string]string)
		queue := list.New()
		visited := make(map[string]bool)

		// Place la salle de départ au début de la liste
		queue.PushBack(startRoom)
		visited[startRoom] = true

		capacity, flow, parent = FindTheWay(queue, startRoom, endRoom, capacity, flow, visited, parent)

		// Arrête la boucle si le chemin n'arrive pas jusqu'à la salle de fin
		if _, foundPath := parent[endRoom]; !foundPath {
			break
		}

		path, maxAnts := BuildPath(endRoom, startRoom, parent, capacity, flow)

		// Met à jour la liste des chemin et le déplacement des fourmis
		for i := 0; i < len(path)-1; i++ {
			u, v := path[i], path[i+1]
			flow[u][v] += maxAnts
			flow[v][u] -= maxAnts
		}
		paths = append(paths, path)
	}

	return paths
}
