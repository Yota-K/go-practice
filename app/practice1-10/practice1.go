package main

import (
	"fmt"
	"math/rand"
)

// MEMO: ここで実行するプログラムは常に同じ環境で実行されるので、
// 擬似乱数を返す rand.Intn はいつも同じ数を返す。
func main() {
	fmt.Println("My favorite Number", rand.Intn(10))
}
