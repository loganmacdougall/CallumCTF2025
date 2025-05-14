package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type GameStart struct {
	Port int
}

func run_game(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var data GameStart
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Fprintf(w, "Failed to parse port number\n")
		return
	}

}

func main() {
	http.HandleFunc("/run_game", run_game)
	log.Fatal(http.ListenAndServe(":8082", nil))
}
