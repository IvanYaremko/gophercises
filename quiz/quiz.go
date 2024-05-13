package quiz

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func QuizGo() {
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'qusetion,answer'")
	flag.Parse()

	_ = csvFilename

	file, err := os.Open(*csvFilename)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
	}

	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		exit("Failed to parse the provided CSV file.")
	}
	problems := parseLines(lines)

	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	correct := 0
	for i, problem := range problems {
		fmt.Printf("Problem %d: %s\n", i+1, problem.q)

		answerCh := make(chan string)
		go func() {
			var answser string
			fmt.Scanf("%s\n", &answser)
			answerCh <- answser
		}()

		select {
		case <-timer.C:
			fmt.Printf("Out of time!")
			return
		case answer := <-answerCh:
			if answer == problem.a {
				correct++
			}
		}

	}

	fmt.Printf("You scored %d out of %d", correct, len(problems))
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

type problem struct {
	q string
	a string
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}
