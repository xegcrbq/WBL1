package main

import "fmt"

//Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.
//
//Например:
//abcd — true
//abCdefAaf — false
//aabcd — false

func IsUnique(s string) bool {
	m := make(map[rune]bool)
	for _, r := range []rune(s) {
		if !m[r] {
			m[r] = true
		} else {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("abcd — ", IsUnique("abcd"))
	fmt.Println("abCdefAaf — ", IsUnique("abCdefAaf"))
	fmt.Println("aabcd — ", IsUnique("aabcd"))
}
