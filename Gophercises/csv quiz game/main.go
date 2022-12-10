package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}

type problem struct {
	question string
	answer   string
}

func parseLines(lines [][]string) []problem {
	// make a flat slice with the equivalent size of "lines"
	slice := make([]problem, len(lines))

	// destruct our 2d slice [[8+1, 9], [3+4, 7]] into our slice array
	// which translates to ["problem", "problem"] or [{5+5 10}, {1+3 4}]
	for index, line := range lines {
		slice[index] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}

	// return the newly parsed slice
	return slice
}

func main() {
	// CODE OVERVIEW
	// present the quiz problem
	// accept the user's input
	// check for correctness

	// define our flags we can receive into the program
	// the flat package defines/returns a poiner
	csvFilename := flag.String("csv", "problems.csv", "a csv file with questions and answers")
	timeLimit := flag.Int("limit", 30, "a time limit in seconds to complete the game")

	// after defining all of our flags, parse them/cache them into the program: https://pkg.go.dev/flag
	flag.Parse()

	// placing this after fmt.Parse() will break our program
	// otherFlag := flag.String("tweet", "google", "testing")

	// open our file's contents with the csvFilename pointer
	// use the actual value from csvFilename and not the pointer directly
	// because os.Open requires a value but not a pointer for it's parameter
	file, err := os.Open(*csvFilename)

	// exit if there is an error
	if err != nil {
		exit(fmt.Sprint("Failed to open this CSV file: ", *csvFilename))
	}

	// begin reading/parsing the CSV file, returns a 2d slice
	csvReader := csv.NewReader(file)
	lines, err := csvReader.ReadAll()

	// exit if there is an error
	if err != nil {
		exit("There was an issue reading this file")
	}

	// build the problems with the "problem" struct
	problems := parseLines(lines)

	// start questioning the user with our game
	score := 0
	timer := time.NewTimer(time.Duration(*timeLimit * int(time.Second)))

	for _, prob := range problems {
		fmt.Printf("What is %s? ", prob.question)

		answerChannel := make(chan string)
		go func() {
			var response string
			fmt.Scanf("%s\n", &response)
			answerChannel <- response
		}()

		select {
		case <-timer.C:
			fmt.Printf("\nYou finished with a score of %g percent!\n", float32(score)/float32(len(problems))*100)
			return
		case response := <-answerChannel:
			if response == prob.answer {
				fmt.Println("Correct!")
				score += 1
			} else {
				fmt.Println("Wrong! The answer ", prob.answer)
			}
		}
	}
}
