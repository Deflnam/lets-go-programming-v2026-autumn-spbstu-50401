package main

import (
	"fmt"

	"github.com/Deflnam/task-1/internal/calculator"
)

func main() {
	var a, b int
	var op string

	_, err := fmt.Scan(&a)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err = fmt.Scan(&b)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, err = fmt.Scan(&op)
	if err != nil {
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
