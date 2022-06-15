package main

import "fmt"

//Реализовать бинарный поиск встроенными методами языка.
func binaryS(elem int, arr []int) bool {
	middle := len(arr) / 2   //берем индекс элемента по середине
	if elem == arr[middle] { //если нашли элемент
		return true
	} else if len(arr) == 1 { //если в массиве всего 1 элемент и он не подходит
		return false
	} else if elem > arr[middle] { //если искомый элемент больше элемента по середине
		arr = arr[middle:] //сужаем область поиска до правой половины массива
	} else if elem < arr[middle] { //если искомый элемент меньше элемента по середине
		arr = arr[:middle] //сужаем область поиска до левой половины массива
	}
	return binaryS(elem, arr)
}

func main() {
	arr := make([]int, 100)
	for i := range arr {
		arr[i] = i
	}
	fmt.Println(binaryS(32, arr))
	fmt.Println(binaryS(100, arr))
}
