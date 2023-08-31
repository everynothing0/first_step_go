package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)


	fmt.Println(now.Unix()) //now
	fmt.Println(now.UnixMilli()) //miliseconds
	fmt.Println(now.UnixNano()) //nanoseconds

	fmt.Println(time.Unix(now.Unix(), 0))
	fmt.Println(time.Unix(0, now.UnixNano()))
}

// go run epoch.go
