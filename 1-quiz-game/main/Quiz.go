package main

import (
	"os"
	"encoding/csv"
	"bufio"
	"log"
	"io"
	"fmt"
	"strings"
	"time"
)

type Quiz struct {
	Question string
	Answer   string
}

func main() {
	csvFile, _ := os.Open("1-quiz-game/problems.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	var quizes []Quiz

	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		quizes = append(quizes, Quiz{strings.TrimSpace(line[0]), strings.TrimSpace(line[1])})
	}

	rightAnswerCount := 0
	triedQuestionCount := 0
	wrongAnswers := make([]string, 0)

	timer1 := time.NewTimer(5 * time.Second)

loop:
	for _, quiz := range quizes {
		select {
		case <-timer1.C:
			break loop

		default:
			fmt.Println(quiz.Question + " = ?")

			var userAnswer string
			fmt.Scanf("%s", &userAnswer)

			if strings.TrimSpace(userAnswer) == quiz.Answer {
				rightAnswerCount++
			} else {
				wrongAnswers = append(wrongAnswers, "\nquestion : " + quiz.Question + ", given answer: " + userAnswer + ""+
					", actual answer: "+ quiz.Answer)
			}

			triedQuestionCount++
		}
	}

	fmt.Printf("Total questions : %d, tried %d, right answer: %d, wrong answers: %d",
		len(quizes), triedQuestionCount, rightAnswerCount, triedQuestionCount - rightAnswerCount)

	if len(quizes) != rightAnswerCount {
		fmt.Println("\nwrong answers : ", wrongAnswers)
	}
}

//https://github.com/gophercises/quiz
