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

type action struct { //наследуем action от human
	human
}

func main() {
	action := action{human{"Andrey", 20}}
	fmt.Printf("Type %T Value: %#v\n", action, action)
	fmt.Printf("action.work(): ") //сравниваем работу методов от action и human
	action.work()
	fmt.Printf("action.human.work(): ")
	action.human.work()
}
