package sol

import (
	"errors"
	"strconv"
)

type str struct {
	index int
	value string
}

// Stringdecoder Возвращает строку с цифрами в строку с символом умноженному числу до него.
// Если в строке есть экранирование, записывается число как есть.
// Если в ходе трансформации произошла ошибка, или же данные невалидные, выводится ошибка.
func Stringdecoder(s string) (string, error) {
	if len(s) <= 1 {
		return s, nil
	}
	mas := []str{}
	resultstring := ""
	if s[0] >= '0' && s[0] <= '9' {
		return "", errors.New("строка не может начинатся с цифры")
	}
	for i := 0; i < len(s); i++ {
		if s[i] == '\\' {
			mas = append(mas, str{index: i + 2, value: "-666"})
			i++
		} else if s[i] >= '0' && s[i] <= '9' {
			c := 0
			j := i + 1
			for j < len(s) {
				if s[j] >= '0' && s[j] <= '9' {
					c++
					j++
				} else {
					break
				}
			}
			mas = append(mas, str{index: i, value: s[i:j]})
			i += c
		} else {
			mas = append(mas, str{index: i + 1, value: "1"})
		}

	}
	if len(mas) == 0 {
		return s, nil
	}
	for i := 0; i < len(mas); i++ {

		atoi, _ := strconv.Atoi(mas[i].value)
		if atoi != 1 {
			atoi--
		}
		if mas[i].value == "-666" {
			resultstring += string(s[mas[i].index-1])
		} else {
			for j := 0; j < atoi; j++ {
				resultstring += string(s[mas[i].index-1])
			}
		}
	}
	return resultstring, nil
}

//Наброски до этого, может будет интересно..
//mas := []string{}
//lastindex := -1
//ind := -1
//for i := 0; i < len(s); i++ {
//if s[i] >= '1' && s[i] <= '9' {
//if i != 0 && lastindex != -1 && i-lastindex == 1 {
//mas[ind] = string(s[lastindex]) + string(s[i])
//} else {
//mas = append(mas, string(s[i]))
//ind++
//lastindex = i
//}
//
//}
//fmt.Println(lastindex)
//}
//fmt.Println(mas)
//mas := []string{}
//for i := 0; i < len(s); i++ {
//if s[i] >= '1' && s[i] <= '9' {
//c := 0
//for j := i + 1; j < len(s); j++ {
//if s[j] >= '1' && s[j] <= '9' {
//c++
//} else {
//mas = append(mas, s[i:j])
//fmt.Println(i)
//i += c
//break
//}
//
//}
//}
//}
//fmt.Println(mas)
//return ""
