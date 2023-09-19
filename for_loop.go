package main
import ("fmt")

func main() {
	for i:=0; i < 8; i++ {
		if i == 4 {
			continue
		}
		if i == 5 {
			break
		}
	fmt.Println(i)
	}
}
