package main

import (
	"fmt"
)

func main() {
	v := "☢🌍♠ ♥ ♦ ♣ ♤ ♡ ♢ ♧ 🃏 ∫ ∬ ∭ ∮ ∯ ∰ ∱ ∲ ∳"
	runeV := []rune(v)
	revV := []rune{}
	for i := len(runeV) - 1; i >= 0; i-- {
		revV = append(revV, runeV[i])
	}
	fmt.Printf("%c", revV)
}
