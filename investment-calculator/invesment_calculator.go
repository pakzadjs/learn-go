package main

import (
	"fmt"
	"math"
)

func main() {
	// invesmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5

	const inflationRate float64 = 2.5
	var invesmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("Invesment Amount: ")
	fmt.Scan(&invesmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := invesmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)

	fmt.Println(futureValue) 
	fmt.Println(futureRealValue)
}