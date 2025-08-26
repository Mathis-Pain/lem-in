package main

import (
	"fmt"
	"lem-in/utils"
	"lem-in/utils/builder"
	"os"
)

func main() {
	n, content := utils.FileMaker(os.Args)
	if content == nil {
		return
	}
	defer content.Close()

	AllPath := builder.PathMaker(content)

	fmt.Printf("Fourmis : %v, Chemins : %v\n", n, AllPath)
}
