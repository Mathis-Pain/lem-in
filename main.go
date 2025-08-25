package main

import (
	"fmt"
	"lem-in/utils/builder"
	"lem-in/utils/checks"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR : Program needs two arguments to run")
		return
	}

	file := os.Args[1]
	content, err := os.Open("./examples/" + file)
	if err != nil {
		fmt.Println("ERROR : Error opening file ", err)
		return
	}
	defer content.Close()

	checks.FirstCheck(content)

	checks.LastCheck(builder.ParseFile(content))

}
