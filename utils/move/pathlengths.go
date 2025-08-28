package move

import (
	"lem-in/models"
	"sort"
)

func SortPaths(AllPaths models.PathList) models.PathList {
	sort.Slice(AllPaths, func(i, j int) bool {
		return len(AllPaths[i]) < len(AllPaths[j])
	})
	return AllPaths
}

func PathLength(AllPaths models.PathList, ants int) models.PathList {
	if len(AllPaths) == 1 {
		return AllPaths
	}
	AllPaths = SortPaths(AllPaths)

	var newPaths models.PathList
	for i := 0; i < len(AllPaths); i++ {
		path := AllPaths[i]
		if len(path) < ants {
			newPaths = append(newPaths, path)
		}

		if i == len(AllPaths)-1 && len(newPaths) == 0 {
			newPaths = append(newPaths, AllPaths[0])
		}
	}

	newPaths = SortPaths(newPaths)

	return newPaths
}
