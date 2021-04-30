package main

import "fmt"
import "github.com/myuser/calculator"

func main() {
	total := calculator.Sum(3, 5)
	println(total)
	println("Version: ", calculator.Version)
	fmt.Println("Hello World!")
}
