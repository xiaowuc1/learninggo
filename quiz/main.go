package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problem struct {
	question string
	answer   string
}

func generate_problems(lines [][]string) []problem {
	problems := make([]problem, len(lines))
	for i, line := range lines {
		// lines[i] is already tokenized by space
		problems[i] = problem{
			question: line[0],
			answer:   line[1],
		}
	}
	return problems
}

func main() {
	filename := flag.String("file", "problems.csv", "file to read questions from")
	// needed to actually resolve the values, note that they're pointers
	flag.Parse()
	file, err := os.Open(*filename)
	if err != nil {
		// file doesn't exist
		fmt.Printf("Failed to open %s\n", *filename)
		os.Exit(1)
	}
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
		// malformed file?
		fmt.Printf("Corrupted file data in %s\n", *filename)
		os.Exit(1)
	}
	problems := generate_problems(lines)
	fmt.Println(problems)

	solved := 0
	for i, problem := range problems {
		fmt.Printf("Problem %d: %s\n", i+1, problem.question)
		var guess string
		fmt.Scanf("%s\n", &guess)
		if guess == problem.answer {
			solved += 1
		}
	}
	fmt.Printf("%d out of %d\n", solved, len(problems))
}
