// is used to create a collection of members of different data types, into a siingle variable.

// syntax 
// type struct_name struct {
//	member1 datatype;
//	...
// }

package main
import ("fmt")

type Person struct {
	name string
	age int
	job string
	salary int
}

func main() {
	var pers1 Person
	var pers2 Person

	// Pers1 specification
	pers1.name = "Edge"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000

	// Pers2
	pers2.name = "Eglantine"
	pers2.age = 17
	pers2.job = "Dogsitter"
	pers2.salary = 4200

	// Access and print
	fmt.Println("Name : ", pers1.name)
	fmt.Println("Age : ", pers1.age, "Age Eglantine : ", pers2.age)
	fmt.Println("Job : ", pers2.job)
	fmt.Println("Salary : ", pers1.salary, "Salary of Eglantine : ", pers2.salary)
}
