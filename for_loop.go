// range keyword is used to more easily iterate over an array, slice or map. returns both he index and the value =
// syntax --> for index, value := array|slice|map
// idx stores the index, val stores the value

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

	fmt.Printf("\n")

	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"apple", "orange", "banana"}
	for i := 0; i < len(adj); i++ {
		for j := 0; j <  len(fruits); j++ {
			fmt.Println(adj[i],fruits[j])
		}
	}

	fmt.Printf("\n")

	for idx, val := range fruits {
		fmt.Printf("%v\t%v\n", idx, val)
	/* if we want to omit the indexes (idx stores the index, val stores the value) :
	for _, val := range fruits {
		fmt.Printf("%v\n", val)
	}
	and reverse :
	for idx, _ idx := range fruits {
		fmt.Printf("%v\n", idx)
	} */
	}
}
