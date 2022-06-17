package main

import "fmt"

//10.	Что выведет данная программа и почему?
func update(p *int) {
	b := 2
	p = &b
}

func main() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p)
	update(p)
	fmt.Println(*p) // 1 1 из-за того, что в функции идет работа с копией значкения указателя
}
