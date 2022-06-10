package main

import "fmt"

//Поменять местами два числа без создания временной переменной.
func main() {
	arr := []int{1, 2, 3}
	fmt.Println(arr)
	arr[0] += arr[2]
	arr[2] = arr[0] - arr[2]
	arr[0] -= arr[2]
	fmt.Println(arr)
	arr[0], arr[1] = arr[1], arr[0]
	fmt.Println(arr)
}
