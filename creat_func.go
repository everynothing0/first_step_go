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

// the same thing but useless (maybe im not sure)

func	funct2(i int, u int) (result int) {
	result = i - u
	return
}

func	example(x int, y string) (result int, text1 string) {
	result = x + x
	text1 = y + "world"
	return
}

func	testcount(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return testcount(x + 1)
}

func	fact_recur(x float64) (y float64) {
	if x > 0 {
		y = x * fact_recur(x-1)
	} else {
		y = 1
	}
	return
}

func main() {

	null()
	message()
	fam_name("Ju", 3)
	message()
	null()
	fmt.Println(funct(1, 2))
	fmt.Println(funct2(8, 4))
	null()
	fam_name("corr", 5)
	fam_name("next", 31)
	message()
	null() // two variables a and b for returns
	_, b := example(5, "hello")
	fmt.Println(b)
	a, _ := example(5, "hello")
	fmt.Println(a)
//	a, b := example(5, "hello" --> not need call one more time)
	fmt.Println(a, b)
	null()
	testcount(1) // start at 1
	null()
	fmt.Println(fact_recur(6)) // y! x!
}
