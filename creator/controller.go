package creator

import (
	"os"
	"path"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
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

	caser := cases.Title(language.English)

	file, _ := os.Create(path.Join("controllers", name+".go"))
	file.WriteString(`package controllers

import (
	"github.com/gin-gonic/gin"
)

/*
* Display a listing of the resource.
 */
func Index` + caser.String(name) + `s(c *gin.Context) {
	
}

/*
* Store a newly created resource in storage.
 */
func Store` + caser.String(name) + `(c *gin.Context) {
	
}

/*
* Display the specified resource.
 */
func Show` + caser.String(name) + `(c *gin.Context) {
	
}

/*
* Update the specified resource in storage.
 */
func Update` + caser.String(name) + `(c *gin.Context) {
	
}

/*
* Remove the specified resource from storage.
 */
func Destroy` + caser.String(name) + `(c *gin.Context) {
	
}


`)

	return control
}
