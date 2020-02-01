package main 

import "fmt"

func add(x int, y int) (sum int) {
	sum = x + y
	return
	// this is a 'naked' return and is okay in smaller functions
}

func main() {
	fmt.Println(add(5, 5))
}