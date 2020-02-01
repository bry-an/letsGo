// like for loops, parens not required

package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

// if statement can start with a short statement to execute before the condition
// note that variables declared inside an if short statement are not available outside the if or else blocks
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	if v := math.Pow(n, x); v < lim {
		return lim
	}
	return lim
}