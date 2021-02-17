package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}

func printQuestions(questions [][]string) {
	for qa := range questions {
		q := qa[0]
		a := qa[1]
		fmt.Println("Question:", q, "Answer:", a)
	}
}

func main() {
	csvFileNamePointer := flag.String("csv", "data.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*csvFileNamePointer)
	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s", *csvFileNamePointer))
	}

	r := csv.NewReader(file)
	data, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV data.")
	}

	fmt.Println(data)
}
