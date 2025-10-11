package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"reflect"
	"bufio"
	"strconv"
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

var solution []int

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
	out_ints, _ := ReadBinaryFile("./challenge/lights.out")

	if reflect.DeepEqual(out_ints, solution) {
		send_completion_message()
		send_result(w, "Program uploaded with correct pattern", 0)
	} else {
		send_result(w, "Program uploaded but with incorrect pattern", 2)
	}
}

func ReadBinaryFile(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var nums []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		val, err := strconv.ParseInt(line, 2, 0)
		if err != nil {
			return nil, fmt.Errorf("failed to parse line %q: %w", line, err)
		}
		nums = append(nums, int(val))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return nums, nil
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
	url := "http://webserver:8080/enter_code"

	req, err := http.NewRequest("POST", url, strings.NewReader("Fix_the_L1ghts_without_I_23485"))
	if err != nil {
		log.Fatal(err)
	}

	// Match Godot headers
	req.Header.Set("Content-Type", "text/plain")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Response Body:", string(body))
}

func main() {
	solution, _ = ReadBinaryFile("./src/solution.out")
	http.HandleFunc("/submit_code", submit_code)
	http.HandleFunc("/light_pattern", light_pattern)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
