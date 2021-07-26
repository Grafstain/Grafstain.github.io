package main

import (
	"fmt"
	"math/rand"
)

const secondPerDay = 86400

func main() {
	company := ""
	distance := 62100000
	trip := ""

	for count := 0; count < 10; count++ {
		switch rand.Intn(3) {
		case 0:
			company = "Space Adventures"
		case 1:
			company = "SpaceX"
		case 2:
			company = "Virgin Galactic"
		}

		speed := rand.Intn(15) + 16
		duration := distance / speed / secondPerDay
		price := rand.Intn(36) + 24

		if rand.Intn(2) == 1 {
			trip = "Round-trip"
			price *= 2
		} else {
			trip = "One-way"
		}

		fmt.Printf("%-16v %4v %10v $%4v\n", company, duration, trip, price)
	}
}
