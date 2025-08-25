package main

import "fmt"

type notGood struct {
}

type Good interface {
	someGoodFunc() string
}

func (not notGood) someNotGoodFunc() string {
	return "Not good"
}

type Adapter struct {
	notGood notGood
}

func (Adapter Adapter) someGoodFunc() string {
	return Adapter.notGood.someNotGoodFunc() + " || Хорошая реализует нехорошую"
}

func main() {
	adapter := Adapter{notGood{}}
	fmt.Println(adapter.someGoodFunc())
}
