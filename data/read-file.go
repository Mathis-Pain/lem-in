package data

import (
	"bufio"
	"fmt"
	"os"
)

func ReadExemple(modele string) []string {
	file, err := os.Open(modele)
	if err != nil {
		fmt.Println("Error open file:", err)
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error scan file", err)
		return nil
	}
	return lines
}
