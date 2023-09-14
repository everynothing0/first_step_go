// +=, -=, %=, /=, ...

package main

import ("fmt")

func main() {
	var x = 5
	fmt.Printf("x is %b\n", x) // 00101
	fmt.Printf("3 is %05b\n", 3) // 00011
	x |= 3 // bitwise inclusive OR
	fmt.Printf("x now is %05b\n", x) // 00111
}
