package main

import "fmt"

func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  return
}

func main() {
  fmt.Println(split(17))

  c, python, java := true, false, "no!"

  var i int
  fmt.Println(i, c, python, java)
}
