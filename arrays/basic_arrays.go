package arrays

import "fmt"

// Defining an Array
func defineArray() [10]int {
	var arr [10]int

	return arr
}

// Comparing Arrays
func compArrays() (bool, bool, bool) {
	var arr1 [5]int
	arr2 := [5]int{0}
	arr3 := [...]int{0, 0, 0, 0, 0}
	// arr4 := [9]int{0, 0, 0, 0, 9} // invalid operation: arr1 == arr4
	arr4 := [5]int{0, 0, 0, 0, 9}

	return arr1 == arr2,
		arr1 == arr3,
		arr1 == arr4
}

// Initializing an Array Using Keys
func initArrayWithKeys() (bool, bool, [10]int) {
	var arr1 [10]int
	arr2 := [...]int{9: 1}
	arr3 := [10]int{1, 9: 10, 4: 5}

	fmt.Printf("arr1: %v\narr2: %d\narr3: %v\n-----\n", arr1, arr2, arr3)

	return arr1 == arr2, arr1 == arr3, arr3
}
