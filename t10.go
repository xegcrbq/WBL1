package main

import (
	"fmt"
	"math"
)

//Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
//Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
//Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.

func group(arr []float64, step int) map[int][]float64 {
	result := make(map[int][]float64)
	for _, value := range arr { //группируем значения
		result[int(math.Round(float64(int(value)/step))*float64(step))] = append(result[int(math.Round(float64(int(value)/step))*float64(step))], value)
	}
	return result
}
func main() {
	fmt.Println(group([]float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}, 10))
}
