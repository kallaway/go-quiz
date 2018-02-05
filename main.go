package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var correctCount int

func main() {
	fmt.Println("Welcome")
	// first, see if there is a flag that specified name of the file.
	// if none: use and read problems.csv
	fileFromCLI := flag.String("file", "problems.csv", "a csv file in a format: question, answer (default is: problems.csv) ")

	flag.Parse()
	fmt.Printf("You've chosen this file: %s\n", *fileFromCLI)

	fileToUse := "quiz-data/" + *fileFromCLI

	problems, err := ioutil.ReadFile(fileToUse)

	if err != nil {
		fmt.Println("There has been an error opening the CSV file.")
	}

	problemsString := string(problems)

	qaPairs := strings.Split(problemsString, "\n")
	questionsCount := len(qaPairs)

	var quizDataSlice [][]string

	for i := range qaPairs {
		// fmt.Printf("%v %v\n", i, v)
		qaPair := strings.Split(qaPairs[i], ",")
		quizDataSlice = append(quizDataSlice, qaPair)
	}

	// fmt.Printf("The length of the quizDataSlice is %v", len(quizDataSlice))

	// ask questions and if they are correct - update score. in the end, post score
	for i := range quizDataSlice {
		qa := quizDataSlice[i]
		q := qa[0]
		a := qa[1]
		fmt.Printf("Q%v: %s\n", (i + 1), q)

		var userResponse string
		reader := bufio.NewReader(os.Stdin)
		userResponse, _ = reader.ReadString('\n')

		userResponse = strings.Trim(userResponse, "\n")
		// fmt.Printf("A is %v of type %T, its length is %v\n", a, a, len(a))
		// fmt.Printf("userResponse is %v of type %T, its length is %v\n", userResponse, userResponse, len(userResponse))

		if userResponse == a {
			correctCount++
		}
	}
	// post score
	fmt.Printf("Game over. You've answered %v out of %v questions correctly.\n", correctCount, questionsCount)
}
