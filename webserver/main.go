package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type GlobalState struct {
	Completed [3]bool
	Code      string
}

var global_state = GlobalState{
	Completed: [3]bool{false, false, false},
	Code:      "XXXX-XXXX-XXXX-XXXX",
}

const code1 = "Fix_the_L1ghts_without_I_23485"
const code2 = "Castle_n0t_impossib1e_9457"
const code3 = "Languag3_hop_h0p_cant_stOp"

var funcMap = template.FuncMap{
	"add": func(a, b int) int { return a + b },
	"all": func(l [3]bool) bool {
		for _, value := range l {
			if !value {
				return false
			}
		}
		return true
	},
}

func landing_page(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	t, err := template.New("landingpage.html").
		Funcs(funcMap).
		ParseFiles("./templates/landingpage.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, global_state)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func enter_code(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	code_bytes, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Didn't receive any code", http.StatusBadRequest)
		return
	}

	var code = string(code_bytes)
	var correct = false

	if strings.Compare(code, code1) == 0 {
		correct = true
		global_state.Completed[0] = true
	} else if strings.Compare(code, code2) == 0 {
		correct = true
		global_state.Completed[1] = true
	} else if strings.Compare(code, code3) == 0 {
		correct = true
		global_state.Completed[2] = true
	}

	if correct {
		io.WriteString(w, "Accepted code")
	} else {
		io.WriteString(w, "Rejected code")
	}
}

func prize(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	t, err := template.New("prize.html").
		Funcs(funcMap).
		ParseFiles("./templates/prize.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, global_state)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func code_check(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	c := r.URL.Query().Get("c")
	num, err := strconv.Atoi(c)
	if err != nil || num < 1 || num > 3 {
		io.WriteString(w, "false")
		return
	}

	index := num - 1 // Get the index

	if global_state.Completed[index] {
		io.WriteString(w, "true")
	} else {
		io.WriteString(w, "false")
	}
}

func main() {
	var steamcode_env = os.Getenv("STEAMCODE")
	if len(steamcode_env) > 0 {
		global_state.Code = steamcode_env
	}

	http.HandleFunc("/prize", prize)
	http.HandleFunc("/enter_code", enter_code)
	http.HandleFunc("/code_check", code_check)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			landing_page(w, r)
			return
		}

		fs.ServeHTTP(w, r)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
