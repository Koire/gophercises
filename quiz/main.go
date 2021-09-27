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

func main() {
	shuffle := flag.Bool("shuffle", false, "Used to shuffle the csv file")
	filename := flag.String("filename", "problems.csv", "Specifies the csv file with the questions")
	// timer := flag.Int("timeLimit", 60, "Default timeout for the quiz")
	timeout := time.After(time.Second * 10)
	finish := make(chan bool)

	csvFile, err := os.Open(*filename)
	if err != nil {
		log.Fatalln("Failed to open", err)
		os.Exit(1)
	}
	r := csv.NewReader(csvFile)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	if *shuffle {
		rand.Shuffle(len(records), func(i, j int) { records[i], records[j] = records[j], records[i] })
	}
	var input string
	var totalCorrect int
	go func() {
		for i, record := range records {
			select {
			case <-timeout:
				fmt.Println("Time is up!")
				finish <- true
				return
			default:
				fmt.Println(i, record[0], "?")
				fmt.Scanf("%s", &input)
				fmt.Println(input)
				if strings.ToLower(strings.TrimSpace(input)) == strings.ToLower(record[1]) {
					totalCorrect++
				}
			}
		}
	}()
	<-finish
	fmt.Printf("Total questions %d: Total Correct: %d \n", len(records), totalCorrect)
}
