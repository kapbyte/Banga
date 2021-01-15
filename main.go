package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "A csv file that contains the problem, answer")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		exit(fmt.Sprintf("Error opening file %s", *csvFilename))
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse csv file")
	}
	fmt.Println(lines)
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][] string) []problem {
	
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
