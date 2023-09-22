package main

import "fmt"

func BhavyaAdd(a, b int) int {
	return a + b
}
func  beantsquareOf5() int {
    return 5 * 5
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
	result := beantsquareOf5()
    	fmt.Println("The square of 5 is:", result)
}
