package main

//Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.
import (
	"fmt"
	"math/big"
)

func main() {
	bigIntA := big.NewInt(1 << 54)
	bigIntB := big.NewInt(1 << 37)
	c := big.NewInt(0)
	fmt.Printf("value1: %v\nvalue2: %v\n", bigIntA, bigIntB)
	fmt.Println("Результат умножения: ", c.Mul(bigIntA, bigIntB))
	fmt.Println("Результат деления: ", c.Div(bigIntA, bigIntB))
	fmt.Println("Результат сложения: ", c.Add(bigIntA, bigIntB))
	fmt.Println("Результат вычитания: ", c.Sub(bigIntA, bigIntB))
}
