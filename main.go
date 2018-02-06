package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var correctCount int

func main() {
	// FLAGS
	// first, see if there is a flag that specified name of the file.
	// if none: use and read problems.csv
	fileFromCLI := flag.String("file", "problems.csv", "a csv file in a format: question, answer (default is: problems.csv)\n")
	// flag.Parse()
	timeLimit := flag.String("limit", "30", "a number of seconds for which the quiz runs.")
	flag.Parse()
	fileToUse := "quiz-data/" + *fileFromCLI

	seconds, err := strconv.Atoi(*timeLimit)
	if err != nil {
		exit("Problem converting time from string to int")
	}

	fmt.Printf("You've chosen this file: %s\n", *fileFromCLI)
	fmt.Printf("Welcome. You have %s seconds to complete the quiz. Press any key to start the quiz.\n", *timeLimit)
	// listen to Enter being pressed (listening for a 'newline')
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	// problems, err := ioutil.ReadFile(fileToUse)
	file, err := os.Open(fileToUse)

	if err != nil {
		exit(fmt.Sprintf("Error when opening file with name: %s", *fileFromCLI))
		fmt.Println("There has been an error opening the CSV file.")
	}

	r := csv.NewReader(file)

	lines, err := r.ReadAll()

	if err != nil {
		exit("Failed to parse the file")
	}

	// problemsString := string(problems)

	problems := parseLines(lines)

	gameTimer := time.NewTimer(time.Second * time.Duration(seconds))

	var numCorrect int

	go func() {
		<-gameTimer.C
		fmt.Printf("\nQuiz has ended. You've answered %v out of %v questions correctly.\n", numCorrect, len(problems))

		os.Exit(0)
	}()

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.a {
			numCorrect++
		}
	}

	fmt.Printf("Duration in seconds %v\n", seconds)

	// defer gameTimer.Stop()

	go func() {
		fmt.Printf("Quiz has ended. You've answered %v out of %v questions correctly.\n", numCorrect, len(problems))
		<-gameTimer.C
		os.Exit(0)
	}()

	// post score
	fmt.Printf("Quiz has ended. You've answered %v out of %v questions correctly.\n", numCorrect, len(problems))
}

type problem struct {
	q string
	a string
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: line[0],
			a: strings.TrimSpace(line[1]),
		}
	}

	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
