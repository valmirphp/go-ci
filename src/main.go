package main

import (
	"fmt"
)

func main() {

	result := soma(5, 5)
	message := fmt.Sprintf("Result: %d", result)

	fmt.Println(message)
	// dialog()

}

func dialog() {

	var inputv1 int
	var inputv2 int

	fmt.Print("Enter an integer number v1: ")
	fmt.Scanf("%d", &inputv1)

	fmt.Print("Enter an integer number v2: ")
	fmt.Scanf("%d", &inputv2)

	result := fmt.Sprintf("Result: %d", soma(inputv1, inputv2))

	fmt.Println(result)

}

func soma(a int, b int) int {
	result := a + b
	return result
}
