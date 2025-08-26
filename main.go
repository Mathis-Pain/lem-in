package main

import (
	"fmt"
	"lem-in/utils"
	"lem-in/utils/builder"
	"os"
)

func main() {
	content := utils.FileMaker(os.Args)
	if content == nil {
		return
	}
	defer content.Close()

	fmt.Println(builder.PathMaker(content))

}
