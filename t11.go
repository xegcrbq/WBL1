package main

import (
	"fmt"
	"math/rand"
)

//Реализовать пересечение двух неупорядоченных множеств.

func main() {
	size := 20
	arr1, arr2, arr3 := make([]int, size), make([]int, size), make([]int, 0)
	for i := 0; i < size; i++ {
		arr1[i] = rand.Intn(10)
		arr2[i] = rand.Intn(10)
	}
	fmt.Printf("Arr1:%v\nArr2:%v\n", arr1, arr2)
	for i1, _ := range arr1 { // запускаю цикл по первому массиву
		for i2, _ := range arr2 { //  в цикле новый цикл по второму массиву
			if arr1[i1] == arr2[i2] && arr1[i1] != -1 { // если есть совпадение чисел первого и второго массива, добавляю это число в новый срез
				arr3 = append(arr3, arr1[i1])
				arr1[i1] = -1 //запоминаем уже пересёкшиеся элементы
				arr2[i2] = -1
			}
		}
	}
	fmt.Printf("Val in arr1 & val in arr2:%v", arr3)
}
