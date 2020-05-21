package main

import "fmt"

func main() {
	var arr []string
	fmt.Println("========================")
	fmt.Println(len(arr))
	fmt.Println(cap(arr))
	fmt.Println("========================")

	arr2 := make([]string, 0)
	fmt.Println("========================")
	fmt.Println(len(arr2))
	fmt.Println(cap(arr2))
	fmt.Println("========================")
}
