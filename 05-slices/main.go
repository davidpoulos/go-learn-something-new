package main

import "fmt"

func main() {
	// `Array Literal` because it has a fixed Capacity + Length
	// food := [5]string{
	// 	"pizza",
	// 	"lasagna",
	// 	"pasta",
	// 	"spaghetti",
	// }
	// fmt.Println(cap(food))
	// fmt.Println(food)

	// `Slice Literal` because it does NOT have a fixed Capacity + Length (capacity can increase)
	sliceOfFoods := []string{
		"pizza",
		"lasagna",
		"pasta",
		"spaghetti",
		"chips",
		"taco",
		"pancake",
	}

	// 	A slice does not store any data, it just describes a section of an underlying array.
	// Changing the elements of a slice modifies the corresponding elements of its underlying array.
	// Other slices that share the same underlying array will see those changes.

	// fmt.Println("Initial capacity", cap(sliceOfFoods))

	// // Give me up to sliceOfFoods[4] (0 to 3)
	// sliceOfFoods = sliceOfFoods[:4]
	// fmt.Println(sliceOfFoods)

	// // Give me 1 and greater
	sliceOfFoods = sliceOfFoods[1:]
	fmt.Println(sliceOfFoods)

	println(len(sliceOfFoods))                   // Should be 2 -- How many elements the slice contains (Actual non-nil element count)
	println("Final Capacity", cap(sliceOfFoods)) // Should be 4 -- How many elements in the `underlying array` (Array size)

	// // Give me first 2 elements
	sliceOfFoods = sliceOfFoods[2:]
	fmt.Println(sliceOfFoods)
	fmt.Println(len(sliceOfFoods))
	fmt.Println(cap(sliceOfFoods))

	//[inclusive:exclusive]

}

//Remove deletes an element of a slice
func Remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	s = s[:len(s)-1]
	return s
}
