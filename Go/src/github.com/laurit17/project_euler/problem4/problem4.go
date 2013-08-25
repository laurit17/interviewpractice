package main

import (
	"fmt"
	"github.com/laurit17/go_utils/string_utils"
	"strconv"
)

func main() {
	maxPalindrome := 0
	for first := 990; first*999 > maxPalindrome; first -= 11 {
		// Here we rely on the fact that palindromes must be divisible by 11.
    // That's why we start with 990, biggest 3-digit number divisible by 11.
		for second := 999; first*second > maxPalindrome; second -= 1 {
			product := first * second
			if product > maxPalindrome &&
				string_utils.IsPalindrome(strconv.Itoa(product)) {
				maxPalindrome = product
			}
		}
	}

	fmt.Printf("Max palindrome is: %d\n", maxPalindrome)
}
