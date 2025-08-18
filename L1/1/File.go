package main

import "fmt"

type human struct {
	name    string
	surname string
	number  int
}

func (human human) Walk() {
	fmt.Println("Walking")
}
func (human *human) changeMyName(newName string) {
	human.name = newName
}

func (human *human) changeMySurname(newSurname string) {
	human.surname = newSurname
}

type action struct {
	human
}

func main() {
	action1 := new(action)
	action1.number = 1
	action1.name = "Sacha"
	action1.surname = "Sarov"
	//Так тоже можно объявить объект экшен, но я для разнообразия сделал по сишному
	//action1 := &action{
	//	human: human{
	//		name:    "Sacha",
	//		surname: "Sarov",
	//		number:  1,
	//	},
	//}
	fmt.Println(action1.name, action1.surname)
	action1.changeMyName("Petr")
	action1.changeMySurname("Pertov")
	fmt.Println(action1.name, action1.surname)
	action1.human.Walk()
}
