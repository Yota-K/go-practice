package calculator

var logMessage = "[LOG]"

// 大文字の変数はどこからでもアクセス可
var Version = "1.0"

// 小文字の関数はパッケージの中からしか呼び出せない
func internalSum(number int) int {
    return number - 1
}

// 大文字の関数はどこからでもアクセス可
func Sum(number1, number2 int) int {
    return number1 + number2
}
