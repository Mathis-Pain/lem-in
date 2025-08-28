package move

import (
	"fmt"
	"lem-in/models"
)

// OneStep should be modified to move a single ant on its specific path
func IAOneStep(ant *models.Ant, path []string, finished int, endRoom string) int {
	// Find the ant's current position within its path
	for i, room := range path {
		if ant.Position == endRoom {
			return finished + 1
		}
		if ant.Position == room {
			// Move the ant to the next room in its path
			ant.Position = path[i+1]
			if ant.Position != endRoom {
				fmt.Printf("%s-%s ", ant.Name, ant.Position) // Example output format
			}
			break
		}
	}
	return finished
}

// MoveAnts refactored for asynchronous movement
func IAMoveAnts(allPaths models.PathList, allAnts int) {
	// Slice to hold ants to be released
	var antsToRelease []models.Ant
	for i := 1; i <= allAnts; i++ {
		antsToRelease = append(antsToRelease, CreateAnt(allPaths, i))
	}

	// Map to associate each ant with its path
	antPaths := make(map[string][]string)

	// Slice to hold ants currently in motion
	var activeAnts []*models.Ant
	finished := 0
	pathIndex := 0

	endRoom := allPaths[0][len(allPaths[0])-1]

	for finished < allAnts {
		// Release a new ant if one is available
		if len(antsToRelease) > 0 {
			for i := 0; i < len(allPaths); i++ {
				// Get the next path to release an ant on
				path := allPaths[pathIndex]

				// Get the next ant from the queue
				ant := &antsToRelease[0]
				antsToRelease = antsToRelease[1:]

				// Assign the ant to a path using the map
				antPaths[ant.Name] = path

				// Add the new ant to the active ants slice
				activeAnts = append(activeAnts, ant)

				// Move to the next path in a cyclic manner
				pathIndex = (pathIndex + 1) % len(allPaths)

				if len(antsToRelease) == 0 {
					break
				}
			}
		}

		// Move all active ants one step
		var nextActiveAnts []*models.Ant
		for _, ant := range activeAnts {
			// Get the ant's path from the map
			path := antPaths[ant.Name]

			if ant.Position != endRoom {
				IAOneStep(ant, path, 0, endRoom) // Move the ant
				nextActiveAnts = append(nextActiveAnts, ant)
			} else {
				finished++
			}
		}

		fmt.Println()               // Newline for each step
		activeAnts = nextActiveAnts // Update the slice of active ants
	}
}
