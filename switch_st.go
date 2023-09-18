// switch statement to select one of many code blocks to be executed
// not need break
// syntax : switch expression { case x:
// is evaluated once
// each value of the switch is compared with the values of each case
// if there is a match, the associated block of code is executed
// the default keyword is optional. It specifies some code if there is no case match

package main ()
import ("fmt")

func main() {
	day := 3

	switch day {
	case 1:
		fmt.Println("monday")
	case 2:
		fmt.Println("tuesday")
	case 3:
		fmt.Println("wednesday")
	case 4:
		fmt.Println("thursday")
	default:
		fmt.Println("OK")
	}
}
