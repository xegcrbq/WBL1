package main

//Реализовать собственную функцию sleep.
import (
	"fmt"
	"time"
)

func Sleep(i int) {
	<-time.After(time.Duration(i) * time.Second)
}
func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		Sleep(1)
	}
}
