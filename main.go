package main

import (
	"fmt"

	functions "github.com/rsauri/winp2000_m06_week5/Functions"
)

func main() {
	//Greetings
	functions.Greeting()
	functions.AdvancedGreeting("Rose")
	fmt.Println()

	//Math
	iNum1 := 10
	iNum2 := 3

	//Addition
	iSum := functions.Add(iNum1, iNum2)
	fmt.Println("ADDITION")
	fmt.Printf("The sum of %d and %d is %d\n", iNum1, iNum2, iSum)
	fmt.Println()

	//Division
	iQuotient, iRemainder := functions.Divide(iNum1, iNum2)
	fmt.Println("DIVISION")
	fmt.Printf("Quotient: %d Remainder: %d\n", iQuotient, iRemainder)
}
