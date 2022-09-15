package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Paranauerj/caipirinha/creator"
	"github.com/Paranauerj/caipirinha/runner"
)

func main() {
	defer Recover()
	ReadLine()
}

func ReadLine() {
	args := os.Args[1:]

	if len(args) < 2 {
		panic("Invalid number of parameters")
	}

	switch strings.ToLower(args[0]) {
	case "create":
		creator.Create(args[1:])
	case "run":
		runner.Run(args[1:])
	default:
		panic("Invalid argument")
	}
}

func Recover() {
	rec := recover()
	if rec != nil {
		fmt.Println(rec)
	}
}
