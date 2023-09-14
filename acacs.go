// access elements of a slice
// access a specific slice element by referring to the index number. index start at 0

// to append syntax --> slice_name = append(slice_name, element1, element2, ...)
// append elements to the end with function append()

// append one slice to another slice
// syntax --> slice3 = append(slice1, slice2...)

// change the lenght of a slice
// unlike arrays, is possible to change the l of a slice

// memory efficiency
// func copy() creates a new underlying array with only the requiered elements for the slice, will reduce the memory used for the program
// syntax --> copy(dest, src)

// make() used for slices, maps or channels, allocates memory on the heap & initia & puts 0 or empty strings into values
// unlike new(), make() returns the same type as its argument, slice: the size determines the L

package main

import ("fmt")

func main () {

	fmt.Printf("\n")

	prices := []int{10,20,30}

	fmt.Println(prices[0], prices[1])
	fmt.Println(prices[2])

	fmt.Printf("\n")
	fmt.Printf("-----------------------------------\n")
	fmt.Printf("\n")

	prices[1] = 50
	fmt.Println(prices[0], prices[1]) // change element of a slice
	fmt.Println(prices[2])

	fmt.Printf("\n")
	fmt.Printf("-----------------------------------\n")
	fmt.Printf("\n")

	slice1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("slice1 = %v\n", slice1)
	fmt.Printf("lenght = %d\n", len(slice1))
	fmt.Printf("capacity = %d\n", cap(slice1))
	slice1 = append(slice1, 20, 21) // append 20 & 21
	fmt.Printf("slice1 = %v\n", slice1)
	fmt.Printf("lenght = %d\n", len(slice1))
	fmt.Printf("capacity = %d\n", cap(slice1))

	fmt.Printf("\n")
	fmt.Printf("-----------------------------------\n")
	fmt.Printf("\n")

	slicea := []int{1, 2, 3}
	sliceb := []int{4, 5, 6}
	slicec := append(slicea, sliceb...) // add slice at the end
	fmt.Printf("slicec = %v\n", slicec)
	fmt.Printf("lenght = %d\n", len(slicec))
	fmt.Printf("capacity = %d\n", cap(slicec))

	fmt.Printf("\n")
	fmt.Printf("-----------------------------------\n")
	fmt.Printf("\n")

	arra := [7]int{7, 8, 9, 10, 11, 12 , 13} // an array
	myslicea := arra[1:5] // slice array
	fmt.Printf("myslicea = %v\n", myslicea)
	fmt.Printf("lenght = %d\n", len(myslicea))
	fmt.Printf("capacity = %d\n", cap(myslicea))
	myslicea = arra[1:3] // change lenght by re-slicing the array
	fmt.Printf("myslicea = %v\n", myslicea)
	fmt.Printf("lenght to reclycle = %d\n", len(myslicea))
	fmt.Printf("capacity to reclycle = %d\n", cap(myslicea))
	myslicea = append(myslicea, 20, 31, 42, 53) // change l by appending items
	fmt.Printf("myslicea by add items = %v\n", myslicea)
	fmt.Printf("lenght by add = %d\n", len(myslicea))
	fmt.Printf("capacity by add = %d\n", cap(myslicea))

	fmt.Printf("\n")
	fmt.Printf("-----------------------------------\n")
	fmt.Printf("\n")

	numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
	// original slice
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("lenght = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))
	// copy, inly need numbers
	needed_numbers := numbers[:len(numbers)-10]
	numbers_copy := make([]int, len(needed_numbers))
	copy(numbers_copy, needed_numbers)
	fmt.Printf("numbers_copy = %v\n", numbers_copy)
	fmt.Printf("lenght = %d\n", len(numbers_copy))
	fmt.Printf("capacity = %d\n", cap(numbers_copy))

	fmt.Printf("\n")
}
