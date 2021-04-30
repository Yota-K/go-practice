package main

import (
  _ "fmt"
  "os"
  "strconv"
)

func main() {
	sum, mul := calc(os.Args[1], os.Args[2])
	println("Sum:", sum)
	println("Mul:", mul)
}

func calc(num1 string, num2 string) (sum int, mul int) {
	int1, _ := strconv.Atoi(num1)
	int2, _ := strconv.Atoi(num2)
	sum = int1 + int2
	mul = int1 * int2
	return
}
