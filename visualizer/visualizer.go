package visualizer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func Visualizer() {
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
}
