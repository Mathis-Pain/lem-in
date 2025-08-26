package builder

// FilterPath s'assurent que deux chemins n'ont aucune pièce en commun
func FilterPath(allPaths [][]string) [][]string {
	// Garde la trace de toutes les pièces qui ont été utilisées par un chemin
	usedRooms := make(map[string]bool)

	// Création de la liste filtrée
	var filteredPaths [][]string

	// Itérer sur chaque chemin dans la liste originale.
	for _, path := range allPaths {
		// Initialise tous les chemins comme valides jusqu'à preuve du contraire
		isPathValid := true

		// Vérifier toutes les pièces du chemin actuel en ignorant "start" et "end" (qui sont dans tous les chemins)
		for i := 1; i < len(path)-1; i++ {
			room := path[i]
			// Si la pièce est déjà utilisée dans un autre chemin, on passe au chemin suivant
			if usedRooms[room] {
				isPathValid = false
				break
			}
		}

		// Si la pièce n'est présente dans aucun des autres chemins, on ajoute le chemin en cours à la liste
		if isPathValid {
			filteredPaths = append(filteredPaths, path)
			// On marque toutes les pièces du chemin qu'on vient d'ajouter comme "déjà utilisées"
			for i := 1; i < len(path)-1; i++ {
				room := path[i]
				usedRooms[room] = true
			}
		}
	}

	return filteredPaths
}
