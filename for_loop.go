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

	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"apple", "orange", "banana"}
	for i := 0; i < len(adj); i++ {
		for j := 0; j <  len(fruits); j++ {
			fmt.Println(adj[i],fruits[j])
		}
	}
}
