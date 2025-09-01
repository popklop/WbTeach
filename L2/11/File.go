package main

import (
	"fmt"
	"strings"
)

func main() {
	poiskAnogram([]string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "стол"})
}

func poiskAnogram(inmaswrong []string) {
	mapa := make(map[string][]int)
	inmas := []string{}
	for i := range inmaswrong {
		inmas = append(inmas, strings.ToLower(inmaswrong[i]))
	}
	for i := range inmas {
		mapa[string(qsort([]rune(inmas[i])))] = append(mapa[string(qsort([]rune(inmas[i])))], i)
	}
	for _, v := range mapa {
		l := make([]string, 0)
		if len(v) > 1 {

			for i := 0; i < len(v); i++ {
				l = append(l, inmas[v[i]])
				continue

			}
			fmt.Printf("%s : %s\n", inmas[v[0]], l)
		}

	}

}

func qsort(s []rune) []rune {
	if len(s) <= 1 {
		return (s)
	}
	lmas := []rune{}
	rmas := []rune{}
	mid := s[len(s)/2]
	for i := 0; i < len(s); i++ {
		if s[i] == mid {
			continue
		}
		if s[i] < mid {
			lmas = append(lmas, (s[i]))
		}
		if s[i] > mid {
			rmas = append(rmas, (s[i]))
		}
	}
	lmasorted := qsort(lmas)
	rmasorted := qsort(rmas)
	return append(append(lmasorted, mid), rmasorted...)
}
