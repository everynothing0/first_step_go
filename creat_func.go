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

func main() {
	message()
	fam_name("Ju", 3)
	message()
	fam_name("corr", 5)
	fam_name("next", 31)
	message()
}
