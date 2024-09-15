package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	numQuestions := flag.Int("num", 10, "the number of questions that will be asked")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()

	if *numQuestions <= 0 {
		fmt.Println("Please enter a number greater than 0")
		os.Exit(1)
	}

	problems := getMathQuestions(*numQuestions)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
askingProblems:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %v = ", i+1, p.q)

		// go routine
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Printf("\n\nYour time expired")
			break askingProblems
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}
	}

	fmt.Printf("\nYou scored %d out of %d\n", correct, len(problems))
}

func getMathQuestions(numQuestions int) []problem {
	questionAnswer := make([]problem, numQuestions)
	for i := 0; i < numQuestions; i++ {
		firstNum, secondNum := rand.Intn(9)+1, rand.Intn(9)+1

		operator := rand.Intn(2) // if 0, use addition; if 1, use subtraction

		var question string
		var answer int

		switch operator {
		case 0:
			question = fmt.Sprintf("%d + %d", firstNum, secondNum)
			answer = firstNum + secondNum
		case 1:
			question = fmt.Sprintf("%d - %d", firstNum, secondNum)
			answer = firstNum - secondNum
		}

		questionAnswer[i] = problem{
			q: question,
			a: strconv.Itoa(answer),
		}
	}

	return questionAnswer
}

type problem struct {
	q string
	a string
}
