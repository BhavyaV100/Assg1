package main

import "fmt"

func BhavyaAdd(a, b int) int {
	return a + b
}

func printTPattern(height int) {
	height = 7
	for i := 0; i < height; i++ {
		for j := 0; j < height; j++ {
			if i == 0 || j == height/2 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func main() {
	result := BhavyaAdd(2, 2)
	fmt.Printf("Bhavya function add: %d\n", result)
	printTPattern(7)
}
