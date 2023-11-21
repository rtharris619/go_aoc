package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"

	"github.com/spf13/viper"
)

func viperEnvVariable(key string) string {
	viper.SetConfigFile("../../.env")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		panic("Invalid type assertion")
	}
	return value
}

func get_file_input() map[string]int {

	result := make(map[string]int)

	client := http.Client{}
	req, err := http.NewRequest("GET", "https://adventofcode.com/2022/day/2/input", nil)

	if err != nil {
		panic(err)
	}

	session_env := viperEnvVariable("SESSION")

	cookieSession := http.Cookie{Name: "session", Value: session_env}
	req.AddCookie(&cookieSession)

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan(); i++ {
		var line = scanner.Text()
		var key = strings.ReplaceAll(line, " ", "")

		if value, ok := result[key]; ok {
			if ok {
				result[key] = value + 1
			}
		} else {
			result[key] = 1
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return result
}

func solve_sample() {

	// A = ROCK = X (1)
	// B = PAPER = Y (2)
	// C = SCISSORS = Z (3)
	// 6 POINTS FOR A WIN
	// 3 POINTS FOR A DRAW

	// 9 different combinations

	// A X => For example occurs 22 times...
	// map["AX"] = 22
	// map["AY"] = 33

	sample := make(map[string]int)
	sample["AY"] = 1
	sample["BX"] = 1
	sample["CZ"] = 1

	solve(sample)
}

func solve(input map[string]int) {
	var total int = 0

	for k, v := range input {

		opponent := k[0]
		me := k[1]

		// step 1: points for what i choose

		if me == 'X' {
			total = total + (1 * v)
		}

		if me == 'Y' {
			total = total + (2 * v)
		}

		if me == 'Z' {
			total = total + (3 * v)
		}

		// step 2: points for win, draw, loss (9 combos)

		if opponent == 'A' {
			if me == 'X' { // draw
				total = total + (3 * v)
			}
			if me == 'Y' { // win
				total = total + (6 * v)
			}
			if me == 'Z' { // loss
				total = total + (0 * v)
			}
		}

		if opponent == 'B' {
			if me == 'X' { // loss
				total = total + (0 * v)
			}
			if me == 'Y' { // draw
				total = total + (3 * v)
			}
			if me == 'Z' { // win
				total = total + (6 * v)
			}
		}

		if opponent == 'C' {
			if me == 'X' { // win
				total = total + (6 * v)
			}
			if me == 'Y' { // loss
				total = total + (0 * v)
			}
			if me == 'Z' { // draw
				total = total + (3 * v)
			}
		}
	}

	println("Total Points:", total)
}

func main() {
	//solve_sample()
	var file_input = get_file_input()
	fmt.Println(file_input)

	solve(file_input)
}
