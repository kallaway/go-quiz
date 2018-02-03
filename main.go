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

	for i, v := range qaPairs {
		fmt.Printf("%v %v\n", i, v)
	}

	var quizDataSlice [][]string

	for i, v := range qaPairs {
		fmt.Printf("%v %v\n", i, v)
		qaPair := strings.Split(qaPairs[i], ",")
		quizDataSlice = append(quizDataSlice, qaPair)
	}

	// ask questions and if they are correct - update score. in the end, post score
}
