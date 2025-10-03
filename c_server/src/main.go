package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

type Data struct {
	Code string
}

/*
	result types:

0 - SUCCESS
1 - FAILED TO COMPILE
2 - COMPILED BUT INCORRECT
3 - COMPILED BUT CRASH
4 - INPUT PARSE ERROR
*/

type Result struct {
	Output string
	Result int
}

var solution []byte

func submit_code(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	data, err := parse_data(r.Body)
	if err != nil {
		send_result(w, "Failed to parse received JSON", 4)
		return
	}

	err = os.WriteFile("./challenge/state.c", []byte(data.Code), 0755)
	if err != nil {
		panic(err)
	}

	cmd := exec.Command("make", "--no-print-directory", "-C", "./challenge")
	out, err := cmd.CombinedOutput()
	if err != nil || cmd.ProcessState.ExitCode() != 0 {
		send_result(w, string(out), 1)
		return
	}

	cmd = exec.Command("./challenge/main")
	out, err = cmd.CombinedOutput()
	if err != nil || cmd.ProcessState.ExitCode() != 0 {
		send_result(w, "CRASHED: program was not uploaded", 3)
		return
	}

	save_current_pattern()

	if bytes.Equal(out, solution) {
		send_completion_message()
		send_result(w, "Program uploaded with correct pattern", 0)
	} else {
		send_result(w, "Program uploaded but with incorrect pattern", 2)
	}
}

func save_current_pattern() {
	cmd := exec.Command("./challenge/main")
	out, _ := cmd.CombinedOutput()
	os.WriteFile("./challenge/lights.out", []byte(out), 0755)
}

func light_pattern(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")

	pattern, _ := os.ReadFile("./challenge/lights.out")
	fmt.Fprintf(w, "%s", pattern)
}

func parse_data(reader io.ReadCloser) (Data, error) {
	decoder := json.NewDecoder(reader)
	var data Data
	err := decoder.Decode(&data)

	return data, err
}

func send_result(w http.ResponseWriter, output string, result int) {
	response := Result{
		Output: output,
		Result: result,
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func send_completion_message() {
	http.Post("http://localhost:8080/enter_code", "text/plain", strings.NewReader("Fix_the_L1ghts_without_I_23485"))
}

func main() {
	solution, _ = os.ReadFile("./src/solution.out")
	http.HandleFunc("/submit_code", submit_code)
	http.HandleFunc("/light_pattern", light_pattern)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
