// https://go-tour-jp.appspot.com/flowcontrol/2
package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	// infinite loop
	// for {
	// 	}
	//
	fmt.Println(sum)
}
