package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"lem-in/data"
	"lem-in/models"
	moveants "lem-in/move-ants"
	"lem-in/path"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var savedMoves string
var savedRooms string

func main() {
	// Vérification des arguments
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <fichier> [--viz]")
		return
	}

	exemple := os.Args[1]

	// Parse le fichier (ton code existant)
	scanExemple := data.ReadExemple(exemple)
	file := data.ExtractFile(scanExemple)
	isCorrect, CorrectFile := data.TestFile(file)
	if !isCorrect {
		fmt.Println("Erreur suite au fichier test-file")
		return
	}

	// Construit les chemins (ton code existant)
	graph := path.BuildGraph(CorrectFile)
	paths := path.FindAllPaths(graph, file.Start, file.End)
	paths = path.SelectPathsOptimizedWithAnts(paths, file.NbAnts)

	// Mode visualisation ou terminal ?
	if len(os.Args) > 2 && os.Args[2] == "--viz" {
		// MODE VISUALISATION
		var buf bytes.Buffer
		moveants.MoveAnts(paths, file.NbAnts, &buf)
		savedMoves = buf.String()

		// Prépare les infos des salles avec coordonnées
		savedRooms = prepareRoomsJSON(file)

		// Affiche aussi dans le terminal
		fmt.Print(savedMoves)

		// Lance le serveur
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "visualizer/visualizer.html")
		})
		http.HandleFunc("/api/moves", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(map[string]string{"moves": savedMoves})
		})
		http.HandleFunc("/api/rooms", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte(savedRooms))
		})

		fmt.Println("\n🚀 Ouvre: http://localhost:8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	} else {
		// MODE TERMINAL NORMAL
		moveants.MoveAnts(paths, file.NbAnts, os.Stdout)
	}
}

func prepareRoomsJSON(file models.File) string {
	roomsData := make(map[string]map[string]interface{})

	// Parse les rooms depuis file.Rooms (format: "nom x y")
	for _, roomLine := range file.Rooms {
		parts := strings.Fields(roomLine)
		if len(parts) >= 3 {
			name := parts[0]
			x, errX := strconv.Atoi(parts[1])
			y, errY := strconv.Atoi(parts[2])

			if errX == nil && errY == nil {
				roomType := "normal"
				if name == file.Start {
					roomType = "start"
				} else if name == file.End {
					roomType = "end"
				}

				roomsData[name] = map[string]interface{}{
					"x":    x,
					"y":    y,
					"name": name,
					"type": roomType,
				}
			}
		}
	}

	jsonData, _ := json.Marshal(roomsData)
	return string(jsonData)
}
