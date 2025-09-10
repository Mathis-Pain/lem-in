package path

import (
	"lem-in/models"
	"math"
)

// Création d'un graphique en go
// Permet  à l'algorithme de voir si le chemin actuel mène à une salle déjà occupée
func BuildGraph(roomsMap map[string]*models.Room, startRoom, endRoom string, links []models.Link) (map[string]map[string]int, map[string]map[string]int) {
	// Première partie map[string] : salle d'origine de la fourmi
	// Deuxième partie map[string]int : salle d'arrivée
	// capacity représente le nombre de fourmis qui peuvent emprunter un chemin (empêche des fourmis d'emprunter des chemins qui mèneraient à la même salle en même temps)
	capacity := make(map[string]map[string]int)
	// flow représente le nombre de fourmis qui empruntent actuellement le chemin en question
	flow := make(map[string]map[string]int)

	// Fabrique le graphique en parcourant la liste des salles
	for roomName := range roomsMap {
		if roomName != startRoom && roomName != endRoom {
			// Crée des identifiants pour l'entrée du chemin et sa sortie
			// Empêche les fourmis de se retrouver à plusieurs dans la même salle même quand elles empruntent des chemins différents
			inNode := roomName + "_in"
			outNode := roomName + "_out"
			capacity[inNode] = make(map[string]int)
			capacity[outNode] = make(map[string]int)
			capacity[inNode][outNode] = 1
			flow[inNode] = make(map[string]int)
			flow[outNode] = make(map[string]int)
		} else {
			capacity[roomName] = make(map[string]int)
			flow[roomName] = make(map[string]int)
		}
	}

	capacity = AddLinks(links, capacity, startRoom, endRoom)

	return capacity, flow
}

// Relie les salles entre elles dans le graphique en parcourant la liste des liens
func AddLinks(links []models.Link, capacity map[string]map[string]int, startRoom, endRoom string) map[string]map[string]int {
	// Parcourt les liens pour les rajouter au graphique
	// Comme la startRoom et la endRoom ont une capacité illimitée, on ne les inclue pas dans les limitations
	for _, link := range links {
		from := link.From
		to := link.To
		if from != startRoom && from != endRoom {
			from += "_out"
		}
		if to != startRoom && to != endRoom {
			to += "_in"
		}
		// Sécurité pour éviter les embouteillages involontaires.
		capacity[from][to] = math.MaxInt64
		capacity[to][from] = math.MaxInt64
	}

	return capacity
}
