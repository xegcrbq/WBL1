package main

import "fmt"

//Разработать программу, которая в рантайме способна определить
//тип переменной: int, string, bool, channel из переменной типа interface{}.
func printType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("This is int ", v)
	case string:
		fmt.Println("This is string", v)
	case bool:
		fmt.Println("this is bool", v)
	case chan int:
		fmt.Println("This is channel", v)
	default:
		fmt.Println("this is not a int, string, bool or channel", v)
	}
}
func main() {
	printType("123")
	printType(123)
	printType(12.3)
	printType(false)
	printType(make(chan int))

}
