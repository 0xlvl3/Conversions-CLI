package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	
	// randChoice will influence which question we get.
	randChoice := rand.Intn(3) + 1 // + 1 so we start at index of 1.
	fmt.Println(randChoice)
	
	// Users score.
	var score int = 0
	fmt.Println(score)


	switch randChoice {
	// Time questions
	case 1:
		// Years => Days

		// Generate a random number of years; y = years.
		y := rand.Intn(100) + 1

		// Ask user to convert years to days.
		fmt.Printf("Convert %v years to days: ", y)

		var d int // d = days.
		fmt.Scanln(&d)
		
		// Check equivalence; a = answer.
		a := y * 365	
        fmt.Println(a)

		if d == a {
			fmt.Printf("Correct! %v years is %v days :)", y, d)
			score ++
		} else {
			fmt.Println("Incorrect :(")
			score --
		}	
	case 2:
		// Days => Years

		// Generate a random number of days; d = days.
		d := rand.Intn(3650) + 366

		// Ask user to convert days to years.
		fmt.Printf("Convert %d days to years, two decimal places: ",d)

		// Read the user input as a formatted string to two decimal places.
		var y float64 // y = years
	    fmt.Scanln(&y)
		y = math.Round(y*100) / 100 // Round to two decimal places.
		
		// Check equivalence; a = answer.
		a := float64(d) / 365.0
		a = float64(a)

		// Comparing absolute of answer and year	
		if math.Abs(a-y) < 0.01 {
			fmt.Println("Correct! :)")
			score ++
			fmt.Println("Score is", score, ":D")
		} else {
			fmt.Println("Incorrect :(")
			score --
			fmt.Println("Score is", score, ":D")
		}	
	case 3:
		// Hours => Days
		
		// Generate random amount of days; h = hours.
		h := rand.Intn(100000) + 100
		fmt.Println("Answer format is integer value rounded to 1 decimal place. e.g. 103 ✓, 102.86 ✗")
		fmt.Printf("How many days are in %v hours: ", h)

		// User input to compare; h = hours.
		var d int
		fmt.Scanln(&d)

		// Check equivalence; a = answer
		a := float64(h) / 24 
		fmt.Println(a)
		a = math.Round(a*10)/10 // Round to 1 decimal place
		fmt.Println(a)
		answer := int(a)
		fmt.Println(answer)

		// if block to check user input == d
	    if d == answer{
			fmt.Printf("Correct %v hours are in %v days", h, answer)
		} else {
			fmt.Println("Incorrect")
		}	

	// todo: finish 4
	case 4:
		// Hours => Days
		// h * d
		
		// Generate random amount of hours; h = hours.
		h := rand.Intn(100000) + 100
		fmt.Println("Answer format is integer value rounded to 1 decimal place. e.g. 103 ✓, 102.86 ✗")
		fmt.Printf("How many days are in %v hours: ", h)

		// User input to compare; h = hours.
		var d int
		fmt.Scanln(&d)

		// Check equivalence; a = answer
		a := float64(h) / 24
		a = math.Round(a*10)/10 // Round to 1 decimal place
		answer := int(a)

		// if block to check user input == d
	    if h == answer{
			fmt.Println("Correct")
		} else {
			fmt.Println("Incorrect")
		}	


	case 5:
		// Hours => Minutes

	case 6:
		// Minutes => Hours

	case 7:
		// Seconds => Minutes

	case 8:
		// Minutes => Seconds

	}
}
