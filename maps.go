// maps are used to store data values in key:value pairs.
// each element in a map is a key:value pair.
// a map is an unordered and changeable collection that does not allow duplicates
// You can find it using the len(0 function. the L of a map is the number of its elements
// the default value of a map is nil.
// maps hold references to an underlying hash table.
// multiple ways for creating maps.
// syntax --> var a = map[KeyType]ValueType{key1:value1, key2:value2,...}
// b := map[KeyType]ValueType{key1:value1, key2:value2,...}
// the map key can of any data type for which the equality operator (==) is defined. Booleans, Numbers, strings, Arrays, Pointers, tructs, Interfaces 
// Invalid key types are: Slices, Maps, Functions. These types are invalid because the equality operator (==) is not defined for them
// The map values can be any type.

// syntax --> value = map_name[key]  Access Map element
// syntax --> update and add map elements --> map_name[key] = value
// syntax --> remove element from map --> delete(map_name, key)
// syntax --> ok :=map_name[key] --> check for specific elements in a map. For only check the existence of a certain key, you can use the blank identifier(_) in place of val.

// maps are references to hash tables
// if two map var refer to the same hash table, changing the content of one variable affect the content of the other.

package main
import ("fmt")

func	new_line() {
	fmt.Printf("\n")
}

func	iterate_over() {
	a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	var b []string // define the order
	b = append(b, "one", "two", "three", "four")

	for k, v := range a { // loop with no order
		fmt.Printf("%v : %v, ", k, v)
	}
	fmt.Println()
	for _, element := range b { // loop with the defined order
		fmt.Printf("%v : %v, ", element, a[element])
	}
	new_line()
}

func main() {

	new_line()

	var a = make(map[string]string) // with make()
	var b map[string]string

	fmt.Println(a == nil)
	fmt.Println(b == nil)

	new_line()

	var c = make(map[string]string)
	c["brand"] = "ford"
	c["model"] = "mustang"
	c["year"] = "1964"

	fmt.Printf(c["brand"])
	new_line()

	fmt.Println(c)

	c["year"] = "1970" // updating an element
	c["color"] = "red" // adding an element

	fmt.Println(c)

	new_line()

	delete(c, "year")

	fmt.Println(c)

	new_line()

	var vso = map[string]string{"br": "for", "mod": "mus", "yea": "196", "day": ""}

	val1, ok1 := vso["br"] // check exist key and its value
	val2, ok2 := vso["color"] // checking for no-existing key and its value
	val3, ok3 := vso["day"] // checking for existing kwy and its value
	val5, _ := vso["mus"]
	_, ok4 := vso["mod"] // only checking for existing and not its value

	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(val5, ok3)
	fmt.Println(ok4)

	new_line()

	ref := vso

	fmt.Println(vso)
	fmt.Println(ref)

	ref["yea"] = "197"
	fmt.Println("after change to b:")

	fmt.Println(vso)
	fmt.Println(ref)

	new_line()
	iterate_over()

	new_line()
}
