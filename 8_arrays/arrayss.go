package main

import "fmt"


func main() {
	// arrays
	var nums[5] int

	fmt.Println(nums) // [0 0 0 0 0]

	// in case of numeric values int or float by default array is initialized with zero, incase of boolean its initialized with false and in string with empty spaces

	// adding element in array
	nums[0]=1
	nums[3]=45

	fmt.Println(nums)
	fmt.Println(nums[1])

	var chars[3] string

	fmt.Println(chars)
	
	chars[0] = "sherlock"
	fmt.Println(chars)
}

