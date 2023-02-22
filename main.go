package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	conversionChoice := rand.Intn(1) + 1 
	// Get num you want (1) + n
	randChoice := rand.Intn(8) + 1 // + 1 so we start at index of 1.

	switch conversionChoice{
	case 1:
		fmt.Println("You've been given time conversions.")
		time_list(randChoice)
	}
}
