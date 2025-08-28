package main

import (
	myntp "8/ntpconsumer"
	"fmt"
	"os"
)

func main() {
	time, err := (myntp.GetTime())
	if err != nil {
		os.Stderr.WriteString(err.Error() + "\n")
		os.Exit(-1)
	}
	fmt.Println(time)
}
