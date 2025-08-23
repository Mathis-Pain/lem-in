package main

import (
	"fmt"
	"lem-in/utils/checks"
	"os"
)

func main() {
	file := os.Args[1]
	content, err := os.Open("./examples/" + file)
	if err != nil {
		fmt.Println("ERROR : Error opening file ", err)
	}
	defer content.Close()

	checks.FirstCheck(content)
}
