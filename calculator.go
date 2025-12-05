package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("Welcome to your Personal Calculator! Input the two operands!\n")
	x := getNumber()
	y:= getNumber()

	for {
		fmt.Printf("Enter the desired operation.\n")
		operation:= getString()

		if len(operation) == 1 && operation == "+" || operation == "-" || operation == "*" || operation == "/"{
			switch operation {
				case "+":
					result := addition(x, y)
					fmt.Printf("%d %v %d = %d\n", x, operation, y, result)
				case "-":
					result := subtraction(x, y)
					fmt.Printf("%d %v %d = %d\n", x, operation, y, result)
				case "*":
					result := multiplication(x ,y)
					fmt.Printf("%d %v %d = %d\n", x, operation, y, result)
				case "/":

					if y != 0 {
						result := division(x , y)
						fmt.Printf("%d %v %d = %.2f\n", x, operation, y, result)
					}
					fmt.Printf("Division by 0 is undefined.")
			}
			break
		}
		fmt.Printf("Invalid input. Please enter one of +, -, *, /\n")
	}
	os.Exit(0)
}

func getString () string {
	var x string
	for {
		fmt.Printf("Scanning for input: ")
		fmt.Scanln(&x)
		if x != "" {
			return x
		}
		fmt.Printf("Incorrect input. Try again.\n")
	}
}

func getNumber () int {
	var x string
	for {
		fmt.Printf("Scanning for input: ")
		fmt.Scanln(&x)
		if x != "" {
			xAtoi, err := strconv.Atoi(x)
			if err == nil {
				return xAtoi
			}
		}
		fmt.Printf("Incorrect input. Try again.\n")
	}
}

func addition (a int, b int) int {
	return a+b
}

func subtraction (a int, b int) int {
	return a-b
}

func multiplication (a int, b int) int {
	return a*b
}

func division (a int, b int) float32 {
	return float32(a)/float32(b)
}
