package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	var (
		A         int
		B         int
		C         int
		c         bool
		i         bool
		v         bool
		F         bool
		n         bool
		forSearch string
		cCounter  int
		re        *regexp.Regexp
	)

	flag.IntVar(&A, "A", 0, "строк после совпадения")
	flag.IntVar(&B, "B", 0, "строк до совпадения")
	flag.IntVar(&C, "C", 0, "строк вокруг совпадения")
	flag.BoolVar(&c, "c", false, "только количество совпадений")
	flag.BoolVar(&i, "i", false, "игнорировать регистр")
	flag.BoolVar(&v, "v", false, "инвертировать совпадение")
	flag.BoolVar(&F, "F", false, "точное совпадение строки (не regexp)")
	flag.BoolVar(&n, "n", false, "нумеровать строки")
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "не указан шаблон поиска")
		os.Exit(1)
	}
	forSearch = flag.Arg(0)
	if C > 0 {
		A += C
		B += C
	}

	if !F {
		pattern := forSearch
		if i {
			pattern = "(?i)" + pattern
		}
		var err error
		re, err = regexp.Compile(pattern)
		if err != nil {
			fmt.Fprintf(os.Stderr, "регулярное выражение не валидно: %v\n", err)
			os.Exit(1)
		}
	} else if i {
		forSearch = strings.ToLower(forSearch)
	}

	var scanner *bufio.Scanner
	if flag.NArg() > 1 {
		file, err := os.Open(flag.Arg(1))
		if err != nil {
			fmt.Fprintf(os.Stderr, "ошибка открытия файла: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()
		scanner = bufio.NewScanner(file)

	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	strnum := 0
	printEnd := 0
	var printStart int
	type bufLine struct {
		num  int
		text string
	}
	buf := []bufLine{}
	printed := make(map[int]bool)

	for scanner.Scan() {
		strnum++
		line := scanner.Text()
		data := line

		if i && F {
			data = strings.ToLower(data)
		}
		if B > 0 {
			buf = append(buf, bufLine{strnum, line})
			if len(buf) > B {
				buf = buf[1:]
			}
		}

		var contains bool
		if F {
			contains = strings.Contains(data, forSearch)
		} else {
			contains = re.MatchString(line)
		}
		if v {
			contains = !contains
		}

		if contains {
			cCounter++
			if B > 0 {
				printStart = strnum - B
				if printStart < 1 {
					printStart = 1
				}
				for _, l := range buf {
					if l.num >= printStart && !printed[l.num] {
						if !c {
							if n {
								fmt.Printf("%d:%s\n", l.num, l.text)
							} else {
								fmt.Println(l.text)
							}
						}
						printed[l.num] = true
					}
				}
			}
			if A > 0 {
				if strnum+A > printEnd {
					printEnd = strnum + A
				}
			}
		}
		if (contains || strnum <= printEnd) && !printed[strnum] {
			if !c {
				if n {
					fmt.Printf("%d:%s\n", strnum, line)
				} else {
					fmt.Println(line)
				}
			}
			printed[strnum] = true
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "ошибка чтения: %v\n", err)
		os.Exit(1)
	}
	if c {
		fmt.Println(cCounter)
	}
}
