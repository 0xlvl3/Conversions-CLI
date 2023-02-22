package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Basic conversions that everyone should know.
---

1. Length/ Distance
- 1 inch = 2.54 centimeters
- 1 foot = 0.3048 meters
- 1 mile = 1.609 kilometers

2. Weight/ Mass
- 1 ounce = 28.35 grams
- 1 pound = 0.454 kilograms
- 1 ton = 907.185 kilograms

3. Temperature
- Celsius to Fehrenheit: F = (C x 1.8) + 32
- Fahrenheit to Celsius: C = (F - 32) / 1.8

4. Volume
- 1 liter = 1000 milliliters
- 1 liter = 1.0567 quarts
- 1 gallon = 3.7854 liters

5. Data storage
- 1 byte = 8bits
- 1 kilobyte (KB) = 1024 Bytes
- 1 megabyte (MB) = 1024 kilobytes
- 1 gigabyte (GB) = 1024 megabytes
- 1 terabyte (TB) = 1024 gigbytes

6. Processing speed
- 1 hertz (Hz) = 1 cycle per second
- 1 megahertz (MHz) = 1 million hertz
- 1 gigahertz (GHz) = 1 billion hertz

7. Energy
- 1 joule (J) = 0.239 calories (cal)
- 1 kilojoule (kJ) = 0.239 kilocalories (kcal)
- 1 calorie (cal) = 4.184 (J)

8. Time 
- 1 year = 365 days
- 1 day = 24 hours
- 1 hour = 60 mins
- 1 minute = 60 seconds


*/

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
