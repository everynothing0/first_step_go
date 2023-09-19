// use the func keyword.
// syntax --> func Function_Name() { code }

package main

import ("fmt")

func message() {
	fmt.Println("I just got executed!")
}

func fam_name(fname string, age int) {
	fmt.Println("Hello", fname, "Ref", age, "year old")
}

func null() {
	fmt.Printf("\n")
}

func funct(x int, y int) int {
	return x + y
}

func main() {

	null()
	message()
	fam_name("Ju", 3)
	message()
	null()
	fmt.Println(funct(1, 2))
	null()
	fam_name("corr", 5)
	fam_name("next", 31)
	message()
	null()
}
