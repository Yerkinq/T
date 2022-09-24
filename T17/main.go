package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	secret := getRandomNumber()
	fmt.Println("Hi! I'm thinking of a number between 1 and 100\nWhat number am I thinking of?")

	for matching := false; !matching; {
		guess := getUserInput()
		matching = isMatching(secret, guess)
	}
}

func isMatching(secret, guess int) bool {
	if guess > secret {
		fmt.Println("Your number is higher than what I had in mind")
		return false
	} else if guess < secret {
		fmt.Println("Your number is lower than what I had in mind")
		return false
	} else {
		fmt.Println("Congrats, you guessed the number!")
		return true
	}
}

func getUserInput() int {
	fmt.Print("Please input your guess: ")
	var input int
	_, err := fmt.Scan(&input)
	if err != nil {
		fmt.Println("Please enter a number!")
	} else {
		fmt.Println("You guessed: ", input)
	}
	return input

}

func getRandomNumber() int {
	rand.Seed(time.Now().Unix())
	return rand.Int() % 101
}
