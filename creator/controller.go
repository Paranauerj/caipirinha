package creator

import (
	"os"
	"path"
	"strings"
)

type Controller struct {
	Path string
	Name string
}

func (c *Controller) Exists() bool {
	folderPath, _ := os.Getwd()
	controllerPath := path.Join(folderPath, "controllers")

	_, err := os.Stat(controllerPath)

	if err != nil {
		return false
	}

	c.Path = controllerPath
	return true
}

func NewController(name string) *Controller {
	control := &Controller{Name: name}
	if !control.Exists() {
		panic("Controllers folder not found!")
	}

	file, _ := os.Create(path.Join("controllers", name+".go"))
	file.WriteString(`package controllers

import (
	"github.com/gin-gonic/gin"
)

/*
* Display a listing of the resource.
 */
func index` + strings.Title(name) + `s(c *gin.Context) {
	
}

/*
* Store a newly created resource in storage.
 */
func store` + strings.Title(name) + `(c *gin.Context) {
	
}

/*
* Display the specified resource.
 */
func show` + strings.Title(name) + `(c *gin.Context) {
	
}

/*
* Update the specified resource in storage.
 */
func update` + strings.Title(name) + `(c *gin.Context) {
	
}

/*
* Remove the specified resource from storage.
 */
func destroy` + strings.Title(name) + `(c *gin.Context) {
	
}


`)

	return control
}
