package main

import "fmt"

//14.	Что выведет данная программа и почему?
//[b b a][a a]
//причины аналогичны t013
func main() {
	slice := []string{"a", "a"}

	func(slice []string) {
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)
	fmt.Print(slice)
}
