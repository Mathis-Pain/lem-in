package main

import (
	"fmt"
	"os"
)

func main() {
	file := os.Args[1]
	content, err := os.ReadFile("./examples/" + file)
	if err != nil {
		fmt.Println("Erreur à l'ouverture du fichier : ", err)
	}

	fmt.Println(string(content))
}
