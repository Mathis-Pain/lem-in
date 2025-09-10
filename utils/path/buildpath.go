package path

import (
	"math"
	"strings"
)

func BuildPath(endRoom, startRoom string, parent map[string]string, capacity, flow map[string]map[string]int) ([]string, int) {
	// Construit le chemin et compte combien de fourmis peuvent être envoyées dedans avant de devoir ouvrir un nouveau chemin
	var path []string
	bottleneck := math.MaxInt64
	for currentRoom := endRoom; currentRoom != startRoom; {
		next := parent[currentRoom]

		// Vérifie les chemins dans les deux sens
		currentCapacity := capacity[next][currentRoom] - flow[next][currentRoom]
		if flow[currentRoom][next] > 0 {
			currentCapacity = flow[currentRoom][next]
		}

		if currentCapacity < bottleneck {
			bottleneck = currentCapacity
		}

		// Retire le "_in" ou le "_out" du nom de la salle avant de l'ajouter au chemin
		roomToAdd := currentRoom
		roomToAdd = strings.TrimSuffix(roomToAdd, "_in")
		roomToAdd = strings.TrimSuffix(roomToAdd, "_out")
		path = append(path, roomToAdd)

		// Passe à la salle suivante
		currentRoom = next
	}
	// Rajoute la salle de départ au chemin
	path = append(path, startRoom)

	// Remet le chemin dans le bon sens (start > end)
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path, bottleneck
}
