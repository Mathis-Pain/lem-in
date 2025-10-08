package main

import (
	"fmt"

	visualizer "lem-in/AI-visualizer"
	"lem-in/data"
	moveants "lem-in/move-ants"
	"lem-in/path"

	"os"
)

func main() {
	// Vérification des arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <fichier> [--viz]")
		return
	}

	exemple := os.Args[1]

	// Parse le fichier
	scanExemple := data.ReadExemple(exemple)
	file := data.ExtractFile(scanExemple)
	isCorrect, CorrectFile := data.TestFile(file)
	if !isCorrect {
		fmt.Println("Erreur suite au fichier test-file")
		return
	}

	// Construit les chemins
	graph := path.BuildGraph(CorrectFile)
	paths := path.FindAllPaths(graph, file.Start, file.End)
	paths = path.SelectPathsOptimizedWithAnts(paths, file.NbAnts)

	// Mode visualisation ou terminal ?
	if len(os.Args) > 2 && os.Args[2] == "--viz" {
		visualizer.Visualizer(paths, file.NbAnts, file)
	} else {
		// MODE TERMINAL NORMAL
		moveants.MoveAnts(paths, file.NbAnts, os.Stdout)
	}
}
