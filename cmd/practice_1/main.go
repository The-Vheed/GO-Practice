package main

import "fmt"

func main() {
	fmt.Println("A simple program to multiply variables")

	var Num int16
	fmt.Print("Enter a number for A: ")
	fmt.Scanln(&Num)

	var Num2 int16
	fmt.Print("Enter a number for B: ")
	fmt.Scanln(&Num2)

	var result int64 = int64(Num) * int64(Num2)
	fmt.Println(Num, "*", Num2, "=", result)
}
