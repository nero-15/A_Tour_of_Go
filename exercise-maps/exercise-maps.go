// https://go-tour-jp.appspot.com/moretypes/23
package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	array := strings.Fields(s)
	result := make(map[string]int)
	for _, w := range array {
		result[w]++
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
