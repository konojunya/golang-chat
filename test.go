package main

import (
	"fmt"
)

func main() {
	// var addr = flag.String("addr",":8080","http service address")

	fmt.Println(add("3", "4"))
}

func add(a, b string) string {
	return a + b
}
