package main

import "fmt"

//12.	Что выведет данная программа и почему?
//0 из-за знака присваивания создастся локальная переменная, с которой будут производиться изменения
func main() {
	n := 0
	if true {
		n := 1
		n++
	}
	fmt.Println(n)
}
