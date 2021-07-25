package creator

import (
	"os"
	"path"
)

type Middleware struct {
	Path string
	Name string
}

func (m *Middleware) Exists() bool {
	folderPath, _ := os.Getwd()
	middlewarePath := path.Join(folderPath, "middlewares")

	_, err := os.Stat(middlewarePath)

	if err != nil {
		return false
	}

	m.Path = middlewarePath
	return true
}

func NewMiddleware(name string) *Middleware {
	middle := &Middleware{Name: name}
	if !middle.Exists() {
		panic("Middlewares folder not found!")
	}

	file, _ := os.Create(path.Join("middlewares", name+".go"))
	file.WriteString(`package middlewares

import (
	"github.com/gin-gonic/gin"
)

func ` + name + `() gin.HandlerFunc {

	return func(c *gin.Context) {

		c.Next()

	}

}
`)

	return middle
}
