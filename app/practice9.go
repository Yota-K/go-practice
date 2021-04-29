package main

import (
	"fmt"
	"strconv"
)

func main() {
	// i, _ := strconv.Atoi("-42")
	// s := strconv.Itoa(-42)
	// fmt.Println(i, s)
	// Goでは変数の値を使用しない時は、_を使用する必要がある
	i, _ := strconv.Atoi("100")
	s := strconv.Itoa(100)
	fmt.Println(i, s)
}
