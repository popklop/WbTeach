package main

import "fmt"

func main() {
	s := "snow dog sun"
	runeS := []rune(s)
	reverseString(runeS)
	fmt.Println(string(runeS))
	reverseWord(runeS)
	fmt.Println(string(runeS))
}

func reverseString(s []rune) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

}

func reverseWord(s []rune) {
	counter := 0
	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] == ' ' {
			for i, j := counter, i-1; i < j; i, j = i+1, j-1 {
				s[i], s[j] = s[j], s[i]
			}
			counter = i + 1
		}
	}
}
