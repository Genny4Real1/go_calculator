package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main () {

// init_calc
	x, err1, y, err2 := init_calc()
	if err1 != nil || err2 != nil{
		fmt.Printf("Error with Calculator Initialization. Exiting program...\n")
		os.Exit(1)
	}

	for {
		fmt.Printf("Enter the desired operation.\n")
		operation:= get_input(false)

		if len(operation) == 1 {
			switch operation {
				case "+":
					result := addition(x, y)
					fmt.Printf("%d %v %d = %d\n", x, operation, y, result)
					os.Exit(0)
				case "-":
					result := subtraction(x, y)
					fmt.Printf("%d %v %d = %d\n", x, operation, y, result)
					os.Exit(0)
				case "*":
					result := multiplication(x ,y)
					fmt.Printf("%d %v %d = %d\n", x, operation, y, result)
					os.Exit(0)
				case "/":
					result := float32(division(x , y))
					fmt.Printf("%d %v %d = %f\n", x, operation, y, result)
					os.Exit(0)
				default:
					fmt.Printf("Error with operation switching. Exiting program...\n")
					os.Exit(0)
			}
		}
		fmt.Printf("Invalid input. Please enter one of +, -, *, /\n")
	}

}

func get_input (is_int bool) (string) {
	var x string
	for {
		fmt.Printf("Scanning for input: ")
		fmt.Scanln(&x)
		if x != "" {
			if is_int == true {
				if check_if_int(x) == nil {
					return x
				}
				continue
			}
			return x
		}
		fmt.Printf("Incorrect input. Try again.\n")
	}
}

func check_if_int (input string) (error) {
	for _, char := range input {
		if unicode.IsDigit(char) == false {
			return errors.New("not valid number")
		}

	}
	return nil
}

func init_calc () (int, error, int, error) { // need to fix returns
	fmt.Printf("Welcome to your Personal Calculator! Input the two operands!\n")
	input1 := get_input(true)
	input2 := get_input(true)

	x, error3 := strconv.Atoi(input1)
	if error3 != nil {
		fmt.Printf("Error with first input conversion to integer value. Exiting Program...\n")
		return 0, errors.New("invalid input"), 0, errors.New("invalid input")
	}

	y, error4 := strconv.Atoi(input2)
	if error4 != nil {
		fmt.Printf("Error with first input conversion to integer value. Exiting Program...\n")
		return 0, errors.New("invalid input"), 0, errors.New("invalid input")
	}
	return x, nil, y, nil
}

func addition (a int, b int) (int) {
	return a+b
}

func subtraction (a int, b int) (int) {
	return a-b
}

func multiplication (a int, b int) (int) {
	return a*b
}

func division (a int, b int) (float32) {
	return float32(a)/float32(b)
}
