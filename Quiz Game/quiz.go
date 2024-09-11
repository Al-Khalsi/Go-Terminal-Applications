package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Reading questions from CSV file
	file, err := os.Open("questions.csv")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

	// Start game 
	fmt.Println("Welcome to the Quiz Game!")
	score := 0

	for _, record := range records {
		question := record[0]
		answer := strings.TrimSpace(record[1]) // Remove extra spaces from the correct answer

		fmt.Println(question)

		// Read the user's response
		reader := bufio.NewReader(os.Stdin)
		userAnswer, _ := reader.ReadString('\n')
		userAnswer = strings.TrimSpace(userAnswer) // Remove extra spaces

		if strings.EqualFold(userAnswer, answer) {
			fmt.Println("Correct!")
			score++
		} else {
			fmt.Println("Wrong! The correct answer is:", answer)
		}
	}

	// Display the final score
	fmt.Printf("Your score: %d out of %d\n", score, len(records))
}