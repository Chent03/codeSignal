package main

import (
	"fmt"
)

func main() {
	inputArray := []int{-23, 4, -3, 8, -12}
	fmt.Println(adjacentElementsProduct(inputArray))
}

func adjacentElementsProduct(inputArray []int) int {
	var (
		greatestProduct int
		current         int
	)
	for i := 0; i < len(inputArray); i++ {
		if i+1 < len(inputArray) {
			if i == 0 {
				greatestProduct = (inputArray[i] * inputArray[i+1])
			}
			current = (inputArray[i] * inputArray[i+1])
			if current > greatestProduct {
				greatestProduct = current
			}
		}
	}
	return greatestProduct
}
