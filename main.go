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

type Bikes struct {
	Name  string
	Type  string
	Owner string
	Price string
}

func (d *Bikes) sent() {
	fmt.Printf("%s is an excellent bike which has %s tires. The owner's name is %s who bought it with %s\n", d.Name, d.Type, d.Owner, d.Price)
}

func sahiltrianglearea(length, width int) int {

	Ar := 0.5 * length * width
	return Ar
}


func tanishaModulus(dividend1, divisor1 int) int {
	dividend1 = 10
	divisor1 = 3 //tanishaModulus is the name of the function
	if divisor1 == 0 {
		return -1 // Handle division by zero error.
	}

	result := dividend1
	for result >= divisor1 {
		result -= divisor1
	}

	return result
	fmt.Printf("The remainder of %d divided by %d is %d\n", dividend1, divisor1, remainder)
}

func  beantsquareOf5() int {
    return 5 * 5
	fmt.Println("The square of 5 is:", result)
}

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


func main() {
	Honda := Bikes{
		Name:  "R15",
		Type:  "2",
		Owner: "Darshan",
		Price: "$125,000",
	}

	Honda.sent()
	result := BhavyaAdd(2, 2)
	fmt.Printf("Bhavya function add: %d\n", result)
	printTPattern(7)
	fmt.Printf("Area of triangle by Sahil : %d", sahiltrianglearea(12, 10))
	EvenOdd()  // call EvenOdd function of Rutvik Pathak
	remainder := tanishaModulus(dividend1, divisor1)
	result := beantsquareOf5()
}
