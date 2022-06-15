package main

import "fmt"

//В какой последовательности будут выведены элементы map[int]int?
func main() {
	m := make(map[int]int)
	m[1] = 2
	m[0] = 0
	m[2] = 4
	fmt.Println(m)
}
