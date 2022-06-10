package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».
func rotateWords(s string) (result string) {
	r := strings.Split(s, " ")
	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-1-i] = r[len(r)-1-i], r[i]
	}
	for _, word := range r {
		result += word + " "
	}
	return
}

func main() {
	fmt.Println(rotateWords("Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode."))
}
