package main

import (
	"errors"
	"fmt"
)

func main() {
	var num1 int
	var num2 int

	fmt.Println("A simple program to divide two integers and return a float result using a function")
	fmt.Print("Enter the two numbers (numerator and denominator) separated by a space: ")
	fmt.Scan(&num1, &num2)

	var result float32
	var err error
	result, err = int_division(num1, num2)
	if err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Printf("%v / %v = %v\n", num1, num2, result)
	}
}

func int_division(num1 int, num2 int) (float32, error) {
	var err error
	if num2 == 0 {
		err = errors.New("division by zero is not allowed")
		return 0, err
	}
	var float_num1 float32 = float32(num1)
	var float_num2 float32 = float32(num2)
	return float_num1 / float_num2, nil
}
