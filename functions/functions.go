// https://go-tour-jp.appspot.com/basics/4
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func kakeru(x int, y int) int {
	return x * y
}

func main() {
	fmt.Println(kakeru(42, 13))
}
