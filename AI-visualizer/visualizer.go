package visualizer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"lem-in/models"
	moveants "lem-in/move-ants"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var savedMoves string
var savedRooms string

func Visualizer(paths [][]string, nbAnts int, file models.File) {
	// Capture les mouvements
	var buf bytes.Buffer
	moveants.MoveAnts(paths, nbAnts, &buf)
	savedMoves = buf.String()

	// Passe aussi paths pour filtrer les salles
	savedRooms = prepareRoomsJSON(file, paths)

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
}

func prepareRoomsJSON(file models.File, paths [][]string) string {
	roomsData := make(map[string]map[string]interface{})
	usedRooms := make(map[string]bool)

	// Marque toutes les salles utilisées dans les chemins
	for _, path := range paths {
		for _, room := range path {
			usedRooms[room] = true
		}
	}
	// Assure-toi que start est inclus
	usedRooms[file.Start] = true

	// Parse uniquement les rooms utilisées
	for _, roomLine := range file.Rooms {
		parts := strings.Fields(roomLine)
		if len(parts) >= 3 {
			name := parts[0]

			// Filtre : n'ajoute que les salles utilisées
			if !usedRooms[name] {
				continue
			}

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
