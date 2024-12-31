package main

import "fmt"

func main() {

	var inputNumber int

	fmt.Println("....<<<< Fibonachi >>>>....")
	fmt.Print("Enter Number : ")
	fmt.Scanln(&inputNumber)

	if inputNumber > 0 {
		avgFibo(inputNumber)
	}

}

func avgFibo(number int) {

	var priveNumber, nextNumber int = 0, 1

	for nextNumber <= number {
		priveNumber = nextNumber
		nextNumber = priveNumber + nextNumber

	}

	fmt.Printf("\n Fibonachi  %v is << %v >>\n", number, nextNumber)
}
