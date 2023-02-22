package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func time_list(n int) {
	// How we get random numbers.
	rand.Seed(time.Now().UnixNano())

	// Users score.
	var score int = 0

	// Helper will be added onto all conversions so they aren't 0. 
	helper := rand.Intn(120) + 1

	switch n {
	// Time questions
	case 1:
		// Years => Days

		// Generate a random number of years; y = Years.
		y := rand.Intn(100) + 1

		// Ask user to convert years to days.
		fmt.Printf("Convert %v years to days: ", y)

		// Get user input; d = Days.
		var d int
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

		// Generate a random number of days; d = Days.
		d := rand.Intn(3650) + helper

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
		// Hours / 24 = Days.
		
		// Generate random amount of days; h = Hours.
		h := rand.Intn(100000) + helper
		fmt.Println("Answer format is integer value rounded to 1 decimal place. e.g. 103 ✓, 102.86 ✗")
		fmt.Printf("How many days are in %v hours: ", h)

		// User input to compare; d = Days.
		var d int
		fmt.Scanln(&d)

		// Check equivalence; a = Answer.
		a := float64(h) / 24 
		a = math.Round(a*10)/10 // Round to 1 decimal place
		answer := int(a)

		// Check user input == answer.
	    if d == answer{
			fmt.Printf("Correct %v hours are in %v days", h, answer)
		} else {
			fmt.Println("Incorrect")
		}	

	case 4:
		// Days => Hours
		// Hours * 24 = Days
		
		// Generate random amount of hours; d = days.
		d := rand.Intn(100000) + helper 
		fmt.Printf("How many hours are in %v days: ", d)

		// User input to compare; h = hours.
		var h int
		fmt.Scanln(&h)

		// Check equivalence; a = answer
		a := d * 24

	    if h == a{
			fmt.Println("Correct")
		} else {
			fmt.Println("Incorrect")
		}	


	case 5:
		// Hours => Minutes
		// 1 => 60

		// Generate random hour; h = Hour
		h := rand.Intn(3750)+ helper 
		fmt.Printf("How many minutes are in %v hours: ",h)

		// Get user answer; m = Minute
		var m int
		fmt.Scanln(&m)

		// a = Answer.
		a := h * 60

		// Compare answer == user input.
		if a == m{
			fmt.Println("Correct")	
		} else {
			fmt.Println("Incorrect")
		}
	

	case 6:
		// Minutes => Hours
		// 60 Mins => 1 Hour

		// Generate random minute; m = Minute.
		m := rand.Intn(100000)+ helper 
		fmt.Printf("How many hours are in %v minutes: ", m)

		// Get user input for hour; h = Hours.
		var h float64
		fmt.Scanln(&h)
		h = math.Round(h*100) / 100 // Round to two decimal places 
	
		// Find answer; a = answer.
		a := float64(m) / 60
		a = float64(a)

		// Compare answer to user input.
		if math.Abs(a-h) < 0.01{
			fmt.Println("Correct!")
		} else {
			fmt.Println("Incorrect")
		}

	case 7:
		// Seconds => Minutes
		// 60 Seconds = 1 Minute.
	
		// Generate random minute; s = Seconds. //Two decimal places.
		s := rand.Intn(1000000)+ helper 
		fmt.Printf("How many minutes are in %v seconds: ", s)

		// Get user input for hour; m = Minutes.
		var m float64
		fmt.Scanln(&m)
		m = math.Round(m*100) / 100 // Round to two decimal places 
	
		// Find answer; a = answer.
		a := float64(s) / 60
		a = float64(a)

		// Compare answer to user input.
		if math.Abs(a-m) < 0.01{
			fmt.Println("Correct!")
		} else {
			fmt.Println("Incorrect")
		}
		

	case 8:
		// Minutes => Seconds
		// 1 Minute => 60 Seconds

		// Generate random minutes; m = Minutes
		m := rand.Intn(5632) + helper 
		fmt.Printf("How many seconds are in %v minutes: ",m)

		// Get user answer; s = Seconds.
		var s int
		fmt.Scanln(&s)

		// a = Answer.
		a := m * 60

		// Compare answer == user input.
		if a == s{
			fmt.Println("Correct")	
		} else {
			fmt.Println("Incorrect")
		}
	}

}
