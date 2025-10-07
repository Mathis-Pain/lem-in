package main

import (
	"fmt"
	"lem-in/data"

	moveants "lem-in/move-ants"

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
	// On construit tous les chemins possible grace aux liaisons
	graph := path.BuildGraph(CorrectFile)
	// On garde tous les chemins trié par le dfs sans room redondante
	paths := path.FindAllPaths(graph, file.Start, file.End)
	// On selectionne les chemins optimaux en fonction du nombre de fourmis
	paths = path.SelectPathsOptimizedWithAnts(paths, file.NbAnts)
	// On les distribue dans les chemins en tour par tour et on affiche le resultat
	moveants.MoveAnts(paths, file.NbAnts)

}
