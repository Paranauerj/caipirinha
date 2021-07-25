package creator

import (
	"strings"
)

func Create(args []string) {
	switch strings.ToLower(args[0]) {
	case "project":
		NewProject(args[1])
	case "middleware":
		NewMiddleware(args[1])
	case "controller":
		NewController(args[1])
	case "model":
		NewModel(args[1])
	default:
		panic("Invalid property to create: " + args[1])
	}
}
