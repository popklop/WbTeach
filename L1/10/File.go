package main

import "fmt"

type floatslice []float64

func main() {
	mapateml := make(map[int]floatslice)
	datamas := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	for i := 0; i < len(datamas); i++ {
		k := (int(datamas[i]) / 10) * 10
		mapateml[k] = append(mapateml[k], datamas[i])
	}
	fmt.Println(mapateml)
}
