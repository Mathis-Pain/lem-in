package main

import (
	"fmt"
	"lem-in/utils"
	"lem-in/utils/builder"
	"lem-in/utils/move"
	"os"
)

func main() {
	ants, content := utils.FileMaker(os.Args)
	if content == nil {
		return
	}
	defer content.Close()

	AllPath := move.PathLength(builder.PathMaker(content), ants)

	if len(AllPath) == 0 {
		fmt.Println("There is no available path.")
		return
	}

	//print.PrintFileData(content)

	fmt.Printf("Fourmis : %v\n", ants)
	for index, r := range AllPath {
		fmt.Printf("Chemin n°%d : %d étapes %v \n", index, len(r), r)
	}

}
