// https://go-tour-jp.appspot.com/flowcontrol/1
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 2; i++ {
		sum += i
	}
	fmt.Println(sum)
}
