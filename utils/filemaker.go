package utils

import (
	"fmt"
	"lem-in/utils/checks"
	"os"
)

func FileMaker(args []string) *os.File {
	var content *os.File

	if len(args) != 2 {
		fmt.Println("ERROR : Program needs two arguments to run")
		return nil
	}
	file := os.Args[1]

	content, err := os.Open("./examples/" + file)
	if err != nil {
		fmt.Println("ERROR : Error opening file ", err)
		return nil
	}

	if !checks.FirstCheck(content) {
		return nil
	}

	return content
}
