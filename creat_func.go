// use the func keyword.
// syntax --> func Function_Name() { code }

package main

import ("fmt")

func message() {
	fmt.Println("I just got executed!")
}

func main() {
	message()
	message()
	message()
}
