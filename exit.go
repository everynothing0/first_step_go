package main

import (
	"fmt"
	"os"
)

func main() {
	
	defer fmt.Println("!") // defer not run when using os, it's mean defering (en deferer)


	os.Exit(5)
}

// go run exit.go
// ./exit
// echo $?
