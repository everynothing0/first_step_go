// syntax --> slice_name := []datatype{values}
// similar to arrays, but + powerful and flexible.

// common way of declaring --> myslice := []int{} 0 lenght 0 capacity
// myslice := []int{1,2,3} 3l 3c

// len() func - returns the lenght of the slice
// cap() func - returns the capacity of the slice

// create a slice from an array

// syntax --> var myarray = [length]datatype{values} // array
// myslice := myarray[start:end] // a slice made from the array

// Create a slice with the make() function
// syntax slice_name := make([]type, length, capacity)
// if the c paramater is not defined, it will be equal to lenght

package main

import ("fmt")

func main () {
	myslice1 := []string{"GO", "slices", "are", "powerful"}

	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	fmt.Printf("\n")
	fmt.Printf("------------------------------------\n")
	fmt.Printf("\n")

	arr1 := [6]int{10, 11, 11, 13, 14, 15}
	myslicearr := arr1[0:4]

	fmt.Printf("mysclicearr = %v\n", myslicearr)
	fmt.Printf("lenght = %d\n", len(myslicearr))
	fmt.Printf("capacity = %d\n", cap(myslicearr))

	fmt.Printf("\n")
	fmt.Printf("------------------------------------\n")
	fmt.Printf("\n")

	myslice2 := make([]int, 5, 10)
	fmt.Printf("myslice2 = %v\n", myslice2)
	fmt.Printf("length = %d\n", len(myslice2))
	fmt.Printf("capacity = %d\n", cap(myslice2))
}
