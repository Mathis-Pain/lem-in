package main

import (
	"fmt"
	"lem-in/utils"
	"lem-in/utils/builder"

	"os"
)

func main() {
	ants, content := utils.FileMaker(os.Args)
	if content == nil {
		return
	}
	defer content.Close()

	AllPath := builder.PathMaker(content)

	if len(AllPath) == 0 {
		fmt.Println("There is no available path.")
		return
	}

	//print.PrintFileData(content)

	fmt.Printf("Fourmis : %v\n", ants)
	for index, r := range AllPath {
		fmt.Printf("Chemin n°%d : %d étapes %v \n", index, len(r), r)
	}

	utils.MoveAnts()
}
