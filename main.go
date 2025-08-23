package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file := os.Args[1]
	content, err := os.Open("./examples/" + file)
	if err != nil {
		fmt.Println("Erreur à l'ouverture du fichier : ", err)
	}
	defer content.Close()

	scanner := bufio.NewScanner(content)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}

func IsValid(b []byte) bool {
	return false
}
