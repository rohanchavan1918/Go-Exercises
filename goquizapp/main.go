package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func exit(msg string) {
	fmt.Println(msg)
	fmt.Println("[!] Exiting ...")
	os.Exit(1)
}

type problem struct {
	// DEclare a problem struct, with the question and answer...
	question string
	answer   string
}

func ParseLines(lines [][]string) []problem {
	// Takes lines read from the csv, and returns a problem struct
	ret := make([]problem, len(lines))
	for i, line := range lines {
		// For every line -> a,b here a is line[0], b is line[1]
		ret[i] = problem{
			question: line[0],
			answer:   line[1],
		}
		// fmt.Println(ret[i])
	}

	return ret

}

func main() {
	// Create a flag to take command line arguments
	csvFileName := flag.String("csv", "problems.csv", "A csv file in the format of questions and answer")
	// csvFileNAme is just a pointer to a string
	flag.Parse()
	file, err := os.Open(*csvFileName)
	if err != nil {
		exit(fmt.Sprintf("[!] Failed to open csv %s", *csvFileName))
	}
	// To read the contents of the csv we need an reader...It reads an IO reader, and the file opened with os.open is already an io.reader, so we can pass directly.

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("[!] Failed to parse CSV")
	}
	// fmt.Println(lines)
	// Pass the lines to parselines func and get the struct
	problems := ParseLines(lines)

	var correct int = 0
	var userAnswer string
	for i, p := range problems {
		fmt.Printf("Problem #%d => %s = \n", i+1, p.question)
		fmt.Scanf("%s\n", &userAnswer)
		if userAnswer == p.answer {
			correct++
		}

	}
	fmt.Printf("Your Score is %d / %d", correct, len(problems))
}
