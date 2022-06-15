package main

//Реализовать паттерн «адаптер» на любом примере.
import "fmt"

type baker struct{}
type smith struct{}

type worker interface {
	work()
}

type smither interface {
	forge()
}

func (b *baker) work() { // действие для пекаря
	fmt.Println(" Baker : baked bread")
}

func (c *smith) forge() { // действие для кузнеца
	fmt.Println(" Smith : forged steel")
}

type workerAdapter struct { //адаптер
	smither smither
}

func newAdapter(s smither) *workerAdapter { //функция инициализации адаптера
	return &workerAdapter{
		smither: s,
	}
}

func (w *workerAdapter) work() { // встраиваем метод, чтобы кузнец мог работать
	w.smither.forge()
}

func main() {
	Billy := &baker{}
	Bob := &smith{}

	Billy.work()
	Bob.forge()
	newAdapter(Bob).work()
}
