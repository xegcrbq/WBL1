package main

import "fmt"

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

type human struct {
	Name string
	Age  int
}

func (h human) work() {
	fmt.Println(h.Name + " is working")
}

type action struct {
	human
}

func main() {
	action := action{human{"Andrey", 20}}
	fmt.Printf("Type %T Value: %#v\n", action, action)
	fmt.Printf("action.work(): ")
	action.work()
	fmt.Printf("action.human.work(): ")
	action.human.work()
}
