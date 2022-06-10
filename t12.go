package main

import "fmt"

type void struct{}

var member void

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]void) // Новое пустое множество
	for _, s := range arr {
		set[s] = member
	}
	for k := range set { // Цикл
		fmt.Println(k)
	}
}
