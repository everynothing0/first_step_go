// don't forget de brackets

package main

import ("fmt")

/*func main() {
	time := 20
	if (time < 18) {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}
} */

func main() {
	a := 14
	b := 14
	if a < b {
		fmt.Println("a is less than b.")
	} else if a > b {
		fmt.Println("a is more than b.")
	} else {
		fmt.Println("a and b are equal.")
	}
}
