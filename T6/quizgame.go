package main

import "fmt"

func main() {
	fmt.Println("Welcome to quiz 1")

	fmt.Printf("Enter your name: ")
	var name string
	fmt.Scan(&name)

	fmt.Printf("Hello, %v, welcome to the game!\n", name)

	fmt.Printf("Enter your age: ")
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Nice, you can play!")
	} else {
		fmt.Println("You cannot play!")
		return
	}

	score := 0
	qs_number := 4

	fmt.Printf("What year did the golang first appear?\n")
	var answer int
	fmt.Scan(&answer)

	if answer == 2009 {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("When go 1.0 was released?\n")
	var answer2 int
	fmt.Scan(&answer2)

	if answer2 == 2012 {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("How many Google Engineers designed Golang?\n")
	var answer3 int
	fmt.Scan(&answer3)

	if answer3 == 3 {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("Which programming language has faster compile time, Go or C++?\n")
	var answer4 string
	fmt.Scan(&answer4)

	if answer4 == "Go" || answer4 == "GO" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect!")
	}

	fmt.Printf("You scored %v out of %v,", score, qs_number)
	percent := (float64(score) / float64(qs_number)) * 100
	fmt.Printf("\nYou scored: %v%%.", percent)

}
