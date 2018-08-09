package main

import (
	"fmt"
)

func main() {
	fmt.Println(shapeArea(5))
}

func shapeArea(n int) int {
	area := 1
	for i := 1; i <= n; i++ {
		area += ((i * 4) - 4)
	}
	return area
}
