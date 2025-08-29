package builder

import (
	"lem-in/models"
	"os"
)

// Explore récursivement les chemins
func explore(current string, end string, links []models.Link, visited map[string]bool, path []string, allPaths *[][]string) {
	// Add the current room to the path.
	path = append(path, current)

	// If the current room is the end room, we've found a complete path.
	if current == end {
		// Make a copy of the path because the slice is reused for backtracking.
		newPath := make([]string, len(path))
		copy(newPath, path)
		*allPaths = append(*allPaths, newPath)
		return
	}

	// Mark the current room as visited to avoid cycles.
	visited[current] = true

	// Explore all links connected to the current room.
	for _, link := range links {
		// Check for a link going from the current room.
		if link.From == current && !visited[link.To] {
			explore(link.To, end, links, visited, path, allPaths)
		} else if link.To == current && !visited[link.From] { // Corrected logic: check the 'From' room for 'visited'
			explore(link.From, end, links, visited, path, allPaths)
		}
	}

	// Backtrack: un-mark the room as visited so it can be part of other paths.
	visited[current] = false
}

func PathMaker(content *os.File) [][]string {
	AllRooms, Links := ParseFile(content)
	var allPaths [][]string
	visited := make(map[string]bool)

	// Lancer l’exploration depuis "start"
	explore(AllRooms.Start.Name, AllRooms.End.Name, Links, visited, []string{}, &allPaths)

	allPaths = FilterPath(allPaths)
	return allPaths
}
