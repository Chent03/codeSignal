package main

import (
	"fmt"
	"strconv"
)

/*
Ticket numbers usually consist of an even number of digits. A ticket number is considered lucky if the sum of the first half of the digits is equal to the sum of the second half.

Given a ticket number n, determine if it's lucky or not.
*/

func main() {
	fmt.Println(isLucky(1230))
}

func isLucky(n int) bool {
	s := strconv.Itoa(n)
	if len(s)%2 != 0 {
		return false
	}
	mid := len(s) / 2
	front := 0
	back := 0

	for num := 0; num < mid; num++ {
		digit := string(s[num])
		bdigit := string(s[(len(s)-1)-num])
		num, err := strconv.Atoi(digit)
		if err != nil {
			return false
		}

		bnum, err := strconv.Atoi(bdigit)
		if err != nil {
			return false
		}
		front += num
		back += bnum
	}

	if front == back {
		return true
	} else {
		return false
	}
}
