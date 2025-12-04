package main

import (
	"errors"
	"fmt"
	"unicode"
	"strconv"
	"os"
)

func main () {

// init_calc
	x, err1, y, err2 := init_calc()
	if err1 != nil || err2 != nil{
		fmt.Printf("Error with Calculator Initialization. Exiting program...")
		os.Exit(1)
	}
	fmt.Printf("%d / %d = %f\n", x, y, float32(x) / float32(y))
	os.Exit(0)
}

func get_input () (string, error) {
	var x string
	fmt.Printf("Scanning for input: ")
	fmt.Scanln(&x)
	if x == "" {
		return "", errors.New("empty input")
	}
	return x, nil
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
	var (
		input1, error1 = get_input()
		input2, error2 = get_input()
	)
	if error1 != nil || error2 != nil {
		fmt.Printf("Invalid Input. Exiting Program...\n")
		return 0, errors.New("invalid input"), 0, errors.New("invalid input")
	}

	if check_if_int(input1) != nil {
		fmt.Printf("First input not a valid number. Exiting Program...\n")
		return 0, errors.New("invalid input"), 0, errors.New("invalid input")
	}
	if check_if_int(input2) != nil {
		fmt.Printf("Second input not a valid number. Exiting Program...\n")
		return 0, errors.New("invalid input"), 0, errors.New("invalid input")
	}

	x, error3 := strconv.Atoi(input1)
	if error3 != nil {
		fmt.Printf("Error with first input conversion to integer value. Exiting Program...")
		return 0, errors.New("invalid input"), 0, errors.New("invalid input")
	}

	y, error4 := strconv.Atoi(input2)
	if error4 != nil {
		fmt.Printf("Error with first input conversion to integer value. Exiting Program...")
		return 0, errors.New("invalid input"), 0, errors.New("invalid input")
	}
	return x, nil, y, nil
}
