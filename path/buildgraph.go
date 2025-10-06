package path

import (
	"lem-in/models"
	"strings"
)

func BuildGraph(file models.File) models.Graph {
	graph := models.Graph{Connection: make(map[string][]string)}

	for _, link := range file.Links {
		parts := strings.Split(link, "-")
		if len(parts) != 2 {
			continue
		}
		a, b := parts[0], parts[1]
		graph.Connection[a] = append(graph.Connection[a], b)
		graph.Connection[b] = append(graph.Connection[b], a)
	}

	return graph
}
