package main

import "fmt"

//Удалить i-ый элемент из слайса.
func removeElem(index int, arr *[]int) {
	if len(*arr) > index {
		arrTemp := *arr
		*arr = append(arrTemp[:index], arrTemp[index+1:]...) //склеиваем 2 части слайса без i-го элемента
	}
}

func main() {
	arr := make([]int, 15)
	for i := range arr {
		arr[i] = i
	}
	fmt.Println(arr)
	removeElem(14, &arr)
	fmt.Println(arr)
}
