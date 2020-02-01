package main

import "fmt"

func forReference() {
	sum := 0
	// note the lack of parens
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	sum = 0
	// the init and post statements are optional:
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}

func whileReference() {
	// For is Go's "while"

	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}