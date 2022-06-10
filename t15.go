package main

import "fmt"

//К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.
//ожидается получить 100 символов, а будет получено 50 из-за кодирования 1 символа 2+ битами
//можно использовать руны, 1 руна = 1 сиивол
var justString string

func createHugeString(size int) (s string) {
	for i := 0; i < size; i++ {
		s += "А"
	}
	return
}

func someFunc() {
	fmt.Printf("%b\n", 1<<10)
	v := createHugeString(1 << 10)
	a := []rune(v)

	justString := string(v[:30])
	justStringRune := string(a[:30])

	fmt.Println(justString)
	fmt.Println(justStringRune)
}

func main() {
	someFunc()
}
