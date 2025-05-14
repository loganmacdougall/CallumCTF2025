package main

import (
	"html/template"
	"log"
	"net/http"
)

type GlobalState struct {
	CompletedFixedLights bool
}

var global_state = GlobalState{CompletedFixedLights: false}

func landing_page(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./templates/landingpage.html")
	t.Execute(w, global_state)
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", landing_page)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
