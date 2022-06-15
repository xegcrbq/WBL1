package main

import (
	"fmt"
	"math/rand"
)

//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1 //находим крайние индексы элементов массива

	pivot := rand.Int() % len(a) //выбираем случайный элемент

	a[pivot], a[right] = a[right], a[pivot] //ставим случайный элемент в конец массива

	for i, _ := range a { //проходимся по всем элементам
		if a[i] < a[right] { //если случайный элемент больше текущего, то ставим текущий в начало массива
			a[left], a[i] = a[i], a[left]

			left++ //сдвигаем индекс начала массива
		}
	}

	a[left], a[right] = a[right], a[left] //меняем местави случайный элемент с элементом после элементов, меньших, чем случайный

	quicksort(a[:left]) //выполняем сортировки для левой и правых частей получившегося массива
	quicksort(a[left+1:])

	return a
}

func main() {
	arr := make([]int, 15)
	for i := range arr {
		arr[i] = rand.Intn(100)
	}
	fmt.Println(arr)
	fmt.Println(quicksort(arr))
}
