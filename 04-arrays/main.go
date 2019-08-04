package main

import "fmt"

func main() {

	// 0 Initialize an array of size 5
	var nums [5]int

	fmt.Println(nums)

	// Initialize with pre-filled values (doesn't have to be equal to size)
	nums2 := [5]int{5, 3, 2, 1}

	fmt.Println(nums2)

	// Access elements

	fmt.Println(nums2[0])

	// Re-assign elements
	nums2[0] = 42

	fmt.Println(nums2[0])

	// Get the length
	fmt.Println(len(nums2))

	fmt.Println(cap(nums2))

}
