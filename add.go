package main 

import "fmt"

func add(x int, y int) (sum int) {
	sum = x + y
	return
}

func main() {
	fmt.Println(add(5, 5))
}