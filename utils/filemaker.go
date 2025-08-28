package utils

import (
	"fmt"
	"lem-in/utils/checks"
	"os"
)

func FileMaker(args []string) (int, *os.File) {
	var content *os.File

	if len(args) != 2 {
		fmt.Println("ERROR <filemaker.go>-l13: Program needs two arguments to run")
		return 0, nil
	}
	file := os.Args[1]

	content, err := os.Open("./examples/" + file)
	if err != nil {
		fmt.Println("ERROR <filemaker.go>-l20: Error opening file ", err)
		return 0, nil
	}

	n, checks := checks.FirstCheck(content)

	if !checks {
		return n, nil
	}

	return n, content
}
