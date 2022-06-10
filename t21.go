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

func (b *baker) work() {
	fmt.Println(" Baker : baked bread")
}

func (c *smith) forge() {
	fmt.Println(" Smith : forged steel")
}

type workerAdapter struct {
	smither smither
}

func newAdapter(s smither) *workerAdapter {
	return &workerAdapter{
		smither: s,
	}
}

func (w *workerAdapter) work() {
	w.smither.forge()
}

func test(a worker) {
	a.work()
}

func main() {
	Billy := &baker{}
	Bob := &smith{}

	workerAdapter := newAdapter(Bob)
	test(Billy)
	test(workerAdapter)
}
