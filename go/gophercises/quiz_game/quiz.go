package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// Problem struct stores the question and the related answer
type Problem struct {
	question string
	answer   string
}

func readArguments() (string, int) {
	filename := flag.String("file", "go/gophercises/quiz_game/problems.csv", "CSV file with the quiz questions")
	timeLimit := flag.Int("time", 30, "Time limit to complete the Quiz")
	flag.Parse()
	return *filename, *timeLimit
}

func setQuizTimer(timeLimit int, timer chan int) {
	time.Sleep(time.Duration(timeLimit) * time.Second)
	timer <- 1
}

func parseCSV(filename string) [][]string {
	csvfile, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	defer csvfile.Close()

	r := csv.NewReader(csvfile)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatalln("Couldn't read the csv file", err)
	}
	return records
}

func problemGenerator(records [][]string) <-chan Problem {
	ch := make(chan Problem, len(records))
	for _, problem := range records {
		ch <- Problem{problem[0], problem[1]}
	}
	close(ch)
	return ch
}

func readInput(done chan string) {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		done <- strings.TrimSpace(scanner.Text())
	}
}

func makeQuestion(problem Problem, timer <-chan int, input <-chan string) (int, error) {
	fmt.Printf("Problem: %v = ", problem.question)
	for {
		select {
		case <-timer:
			return 0, errors.New("timed out")
		case resp := <-input:
			if resp == problem.answer {
				return 1, nil
			}
			return 0, nil
		}
	}
}

func finalScore(score, total int) {
	fmt.Printf("\nYou scored %v out of %v.", score, total)
}

func startQuiz(total int, problem <-chan Problem, timer <-chan int) {
	score := 0
	input := make(chan string)
	go readInput(input)

	for {
		p, ok := <-problem
		if ok == false {
			finalScore(score, total)
			break
		}

		result, err := makeQuestion(p, timer, input)
		if err != nil {
			finalScore(score, total)
			break
		}
		score += result
	}
}

func main() {
	filename, timeLimit := readArguments()
	records := parseCSV(filename)

	problem := problemGenerator(records)
	timer := make(chan int)

	go setQuizTimer(timeLimit, timer)
	startQuiz(len(records), problem, timer)
}
