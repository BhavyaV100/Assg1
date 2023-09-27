package main

import "fmt"

func EvenOdd() {
	var input int
	// Checking number is even or odd
	fmt.Println("Enter the number :")
	fmt.Scanln(&input)

	if input%2 == 0 {
		fmt.Println("The numebr is Even")
	} else {
		fmt.Println("The number is Odd")
	}
}
