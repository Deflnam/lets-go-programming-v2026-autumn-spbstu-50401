package main

import "fmt"

func main() {
	var n int

	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	for i := 0; i < n; i++ {
		var k int
		_, err = fmt.Scan(&k)
		if err != nil {
			fmt.Println("Invalid input")
			return
		}

		low := 15
		high := 30

		for j := 0; j < k; j++ {

		}
	}
}
