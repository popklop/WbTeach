package main

import (
	"fmt"
	"math/big"
)

func umnozh(a big.Int, b big.Int) string {
	res := new(big.Int)
	res.Mul(&a, &b)
	return res.String()
}

func podeli(a big.Int, b big.Int) string {
	res := new(big.Int)
	res.Div(&a, &b)
	return res.String()
}
func slozhi(a big.Int, b big.Int) string {
	res := new(big.Int)
	res.Add(&a, &b)
	return res.String()
}
func vichti(a big.Int, b big.Int) string {
	res := new(big.Int)
	res.Sub(&a, &b)
	return res.String()
}

func main() {
	var operation string
	var a, v big.Int
	fmt.Scan(&a)
	fmt.Scan(&operation)
	fmt.Scan(&v)
	switch operation {
	case "+":
		fmt.Println(slozhi(a, v))
	case "-":
		fmt.Println(vichti(a, v))
	case "/":
		fmt.Println(podeli(a, v))
	case "*":
		fmt.Println(umnozh(a, v))
	default:
		fmt.Println("Такой операции нет")
		return
	}

}
