package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	csvFilename := flag.String("csv", "problems.csv", "A csv file that contains the problem, answer")
	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		fmt.Printf("Error opening file %s", *csvFilename)
		os.Exit(1)
	}
	_ = file
	fmt.Printf("Success opening file %s", *csvFilename)
}
