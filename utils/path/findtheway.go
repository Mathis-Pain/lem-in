package path

import "container/list"

func FindTheWay(queue *list.List, startRoom, endRoom string, capacity, flow map[string]map[string]int, visited map[string]bool, parent map[string]string) (map[string]map[string]int, map[string]map[string]int, map[string]string) {
	// Système BFS (Breadth-First Search), algorithme de recherche dans un graphique
	// Boucle qui continue tant qu'elle n'a pas trouvé un chemin valide entre la salle de départ et la salle de fin
	for queue.Len() > 0 {
		current := queue.Remove(queue.Front()).(string)
		if current == endRoom {
			// Interrompt la boucle si on arrive à la salle finale
			break
		}

		// Vérifie s'il est possible pour la fourmi d'emprunter le chemin vers une des salles suivantes
		for next := range capacity[current] {
			// Vérifie si aucune fourmi n'est dans la prochaine salle et si une fourmi n'est pas déjà en train d'y aller par un autre chemin
			if !visited[next] && capacity[current][next]-flow[current][next] > 0 {
				parent[next] = current
				visited[next] = true
				queue.PushBack(next)
			}
		}
		// Même vérifications en remontant en arrière dans le chemin pour voir s'il n'y a pas un chemin plus efficace
		for previous := range capacity {
			if _, ok := capacity[previous][current]; ok {
				if !visited[previous] && flow[previous][current] > 0 {
					parent[previous] = current
					visited[previous] = true
					queue.PushBack(previous)
				}
			}
		}
	}

	return capacity, flow, parent
}
