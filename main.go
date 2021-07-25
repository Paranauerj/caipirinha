package main

import (
	"fmt"
	"myprojects/caipirinha/creator"
	"myprojects/caipirinha/runner"
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
		creator.Create(args[1:])
	}

	if strings.ToLower(args[0]) == "run" {
		runner.Run(args[1:])
	}
}

func Recover() {
	rec := recover()
	if rec != nil {
		fmt.Println(rec)
	}
}
