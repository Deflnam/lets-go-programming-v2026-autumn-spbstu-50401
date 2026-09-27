package main

import (
	"fmt"

	"github.com/Deflnam/task-1/internal/calculator"
)

func main() {
	var a, b int
	var op string

	n, err := fmt.Scan(&a)
	if err != nil || n != 1 {
		fmt.Println("Invalid first operand")
		return
	}

	n, err = fmt.Scan(&b)
	if err != nil || n != 1 {
		fmt.Println("Invalid second operand")
		return
	}

	n, err = fmt.Scan(&op)
	if err != nil || n != 1 {
		fmt.Println("Invalid operation")
		return
	}

	result, err := calculator.Calculate(a, b, op)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(result)
}
