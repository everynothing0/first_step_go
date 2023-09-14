// access elements of a slice
// access a specific slice element by referring to the index number. index start at 0

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
}
