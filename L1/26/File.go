package main

import "fmt"

func main() {
	s := "abcd"
	fmt.Println(stringtester(s))

}

func stringtester(s string) bool {
	mapa := make(map[byte]struct{})
	var tekushiybyte byte
	for i := 0; i < len(s); i++ {
		if s[i] >= 65 && s[i] <= 90 {
			tekushiybyte = s[i] + 32
		} else {
			tekushiybyte = s[i]
		}
		if _, ok := mapa[tekushiybyte]; ok {
			return false
		}
		mapa[tekushiybyte] = struct{}{}

	}
	return true
}
