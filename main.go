package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Initialize the random number generator.
	rand.Seed(time.Now().UnixNano()) 
	// The seed value is used to generate a sequence of random numbers.
	// Using the same seed value will result in the same random numbers.
	// So we use time.Now and UnixNano() to generate a current timestamp that will 
	// Ensure a new random sequence of numbers is generated everytime.

	// Generate a random number between 100 and 1000 millimeters.
	mm := rand.Intn(901) + 100 // rand.Intn(901) generates 0 - 900
	// We then + 100 onto whatever number that number is.

	// Convert millimeters to meters.
	m := float64(mm) / 1000.0

	// Ask the user to convert millimeters to meters.
	fmt.Printf("Convert %d millimeters to meters: ", mm)

	// Read the user's input.
	var answer float64
	fmt.Scanln(&answer)

	// Compare the user's answer to the correct answer.
	if answer == m {
		fmt.Println("Correct!")
	} else {
		fmt.Printf("Incorrect. The answer is %.2f meters.\n", m)
	}
}
