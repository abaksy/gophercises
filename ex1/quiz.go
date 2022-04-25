package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

/*
Takes the path to the CSV file containing questions and answers as argument.
Parses the CSV file using the `encoding/csv` package and returns an array of records
*/
func readCSVFile(fileName string) [][]string {
	f, err := os.Open(fileName)
	if err != nil { // Error in opening the CSV file
		log.Println("Error in opening file!")
		panic("Error in opening file!")
	}
	defer f.Close()
	fmt.Println("Reading questions from", fileName, "...")

	csvReader := csv.NewReader(f)
	questions, err := csvReader.ReadAll()
	if err != nil { // Error in parsing the CSV file
		log.Println("Error in parsing CSV file!")
		panic("Error in parsing CSV file!")
	}
	return questions
}

/*
While the timer is running, ask questions.
If the timer sends a value through its channel, then the timer has finished running, and execution is interrupted.
Until then, loop over questions and ask them
*/
func runTimedQuiz(questions [][]string, timeLimit int) {
	score := 0
	done := false
	numQuestions := len(questions)
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	for _, value := range questions {
		select {
		case <-timer.C:
			fmt.Println("You scored:", score, "out of", numQuestions)
			done = true
			return
		default:
			question, correctAns := value[0], value[1]
			fmt.Printf("%s: ", question)
			var userAns string
			fmt.Scanln(&userAns)
			userAns = strings.TrimSpace(userAns)
			userAns = strings.ToLower(userAns)
			if strings.Compare(userAns, correctAns) == 0 {
				score += 1
			}
		}
	}
	if !done {
		fmt.Println("You scored:", score, "out of", numQuestions)
	}
}

func main() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666) // Open log file and create if not existing
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	// Set up command line arguments
	var fileNameArg = flag.String("file", "problems.csv", "Filename of the problems list")
	var timerArg = flag.Int("time", 30, "Time limit for quiz")
	var shuffleArg = flag.Bool("shuf", false, "Shuffle questions each time the quiz is run")
	flag.Parse()
	var fileName string = *fileNameArg
	var timeLimit int = *timerArg

	// Read CSV file and get questions
	questions := readCSVFile(fileName)

	//Shuffle questions if flag is set to true
	if *shuffleArg == true {
		rand.Shuffle(len(questions), func(i, j int) {
			questions[i], questions[j] = questions[j], questions[i]
		})
	}
	runTimedQuiz(questions, timeLimit)
}
