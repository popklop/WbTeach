package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println(poiskAnogram([]string{"пятак", "пятка", "тяпка", "атяпк", "листок", "слиток", "столик", "стол"}))
}

func poiskAnogram(inmaswrong []string) map[string][]string {
	mapa := make(map[string][]string)
	inmas := []string{}
	for i := range inmaswrong {
		inmas = append(inmas, strings.ToLower(inmaswrong[i]))
	}
	for i := range inmas {
		mapa[string(qsort([]rune(inmas[i])))] = append(mapa[string(qsort([]rune(inmas[i])))], (inmas[i]))
	}
	for k, v := range mapa {
		if len(v) > 1 {
			val := mapa[k]
			delete(mapa, k)
			mapa[v[0]] = val
			sort.Strings(val)
		} else {
			delete(mapa, k)
		}
	}
	return mapa
}

func qsort(s []rune) []rune {
	if len(s) <= 1 {
		return s
	}
	lmas := []rune{}
	rmas := []rune{}
	mid := s[len(s)/2]
	midarr := []rune{}
	for i := 0; i < len(s); i++ {
		if s[i] == mid {
			midarr = append(midarr, s[i])
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
	return append(append(lmasorted, midarr...), rmasorted...)
}
