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
var savedLinks string

func Visualizer(paths [][]string, nbAnts int, file models.File) {
	// Capture les mouvements
	var buf bytes.Buffer
	moveants.MoveAnts(paths, nbAnts, &buf)
	savedMoves = buf.String()

	// Passe aussi paths pour filtrer les salles
	savedRooms = prepareRoomsJSON(file, paths)
	savedLinks = prepareLinksJSON(paths)

	// Affiche aussi dans le terminal
	fmt.Print(savedMoves)

	// Lance le serveur
	// ✨ 1. D'ABORD les fichiers CSS et JS
	http.HandleFunc("/styles.css", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "AI-visualizer/styles.css")
	})
	http.HandleFunc("/script.js", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "AI-visualizer/script.js")
	})

	// ✨ 2. ENSUITE les routes API
	http.HandleFunc("/api/moves", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"moves": savedMoves})
	})
	http.HandleFunc("/api/rooms", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(savedRooms))
	})
	http.HandleFunc("/api/links", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(savedLinks))
	})

	// ✨ 3. EN DERNIER la route "/" pour le HTML
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "AI-visualizer/visualizer.html")
	})

	fmt.Println("\n Ouvre: http://localhost:8080")
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

func prepareLinksJSON(paths [][]string) string {
	linksSet := make(map[string]bool)
	var links []map[string]string

	// Extrait toutes les connexions des chemins
	for _, path := range paths {
		for i := 0; i < len(path)-1; i++ {
			from := path[i]
			to := path[i+1]

			// Crée une clé unique pour éviter les doublons
			key := from + "-" + to
			reverseKey := to + "-" + from

			if !linksSet[key] && !linksSet[reverseKey] {
				linksSet[key] = true
				links = append(links, map[string]string{
					"from": from,
					"to":   to,
				})
			}
		}
	}

	jsonData, _ := json.Marshal(links)
	return string(jsonData)
}
