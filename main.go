package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	defer Recover()
	ReadLine()
}

func ReadLine() {
	args := os.Args[1:]

	if strings.ToLower(args[0]) == "create" {
		Creator(args[1:])
	}
}

func Recover() {
	rec := recover()
	if rec != nil {
		fmt.Println(rec)
	}
}
