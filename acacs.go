// access elements of a slice
// access a specific slice element by referring to the index number. index start at 0

// to append syntax --> slice_name = append(slice_name, element1, element2, ...)
// append elements to the end with function append()

// append one slice to another slice
// syntax --> slice3 = append(slice1, slice2...)

// change the lenght of a slice
// unlike arrays, is possible to change the l of a slice

package main

import ("fmt")

func main () {
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

	arra := [7]int{7, 8, 9, 10, 11, 12 , 13, 14} // an array
	myslicea := arra[1:5] // slice array
	fmt.Printf("myslicea = %v\n", myslicea)
	fmt.Printf("lenght = %d\n", len(myslicea))
	fmt.Printf("capacity = %d\n", cap(myslicea))
}
