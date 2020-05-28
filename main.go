package main

import (
	"fmt"
	"strings"
)

func main() {
	nums := "1.345.679"
	nums = strings.ReplaceAll(nums, ".", "")

	for i, num := range nums {
		stringRepeat := string(num)
		stringRepeat += strings.Repeat("0", len(nums)-(i+1))
		fmt.Println(stringRepeat)
	}
}
