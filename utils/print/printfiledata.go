package print

import (
	"bufio"
	"fmt"
	"os"
)

// Print le contenu du fichier (nombre de fourmis, salles, tunnels)
func PrintFileData(file *os.File) {
	file.Seek(0, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

	}
}
