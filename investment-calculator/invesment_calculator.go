package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5

func main() {
	// invesmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5
	var invesmentAmount float64
	var years float64
	var expectedReturnRate float64

	// fmt.Print("Invesment Amount: ")
	outputText("Invesment Amount: ")
	fmt.Scan(&invesmentAmount)

	// fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValue(invesmentAmount, expectedReturnRate, years)
	// futureValue := invesmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	// futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)


	// formattedFV := fmt.Sprintf("future Value: %.1f\n", futureValue)
	// formattedRFV := fmt.Sprintf("Future Value (adjusted for inflation): %.1f\n", futureRealValue)
	// fmt.Println("Future Value", futureValue)
	fmt.Printf("future Value: %.1f\nFuture Value (adjusted for inflation): %.1f", futureValue, futureRealValue) 

	// fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(invesmentAmount , expectedReturnRate , years float64) (float64, float64)  {
	fv := invesmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	rfv := fv / math.Pow(1 + inflationRate / 100, years)

 return	fv, rfv
}