package main

import (
	"fmt"
	"lem-in/data"
	"lem-in/filterpath"

	"lem-in/path"

	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Erreur argument programme different de trois")
	}
	exemple := os.Args[1]
	scanExemple := data.ReadExemple(exemple)
	file := data.ExtractFile(scanExemple)
	isCorrect, CorrectFile := data.TestFile(file)
	if !isCorrect {
		fmt.Print("Erreur suite au fichier test-file")
		return
	}
	graph := path.BuildGraph(CorrectFile)
	paths := path.FindAllPaths(graph, file.Start, file.End)
	paths = filterpath.UniquePaths(paths)
	paths = filterpath.EssentialPaths(paths)

	for i, path := range paths {
		fmt.Printf("Chemin %d (%d salles) : %v\n", i+1, len(path), path)
	}
}
