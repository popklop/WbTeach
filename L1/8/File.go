package main

import "fmt"

func main() {
	//true - установка бита в 1, false - установка бита в 0
	fmt.Println(changebyte(5, 0, false))
}

func changebyte(val int64, i int, bit bool) int64 {
	var mask int64 = 1 << i
	if bit {
		return val | mask
	} else {
		return val &^ mask
	}
}
