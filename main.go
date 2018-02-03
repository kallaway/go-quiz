package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var correctCount int

func main() {
	fmt.Println("Welcome")
	// first, see if there is a flag that specified name of the file.
	// if none: use and read problems.csv

	problems, err := ioutil.ReadFile("quiz-data/problems.csv")

	if err != nil {
		fmt.Println("There has been an error opening the CSV file.")
	}

	problemsString := string(problems)

	qaPairs := strings.Split(problemsString, "\n")
	questionsCount := len(qaPairs)

	var quizDataSlice [][]string

	for i, v := range qaPairs {
		fmt.Printf("%v %v\n", i, v)
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
		fmt.Scanln(userResponse)

		if a == userResponse {
			correctCount++
		}
	}
	// post score
	fmt.Printf("Game over. You've answered %v out of %v questions correctly.", correctCount, questionsCount)
}
