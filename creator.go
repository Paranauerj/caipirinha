package main

import (
	"os"
	"path"
	"strings"
)

type Project struct {
	Path string
	Name string
}

type Middleware struct {
	Path string
	Name string
}

type Controller struct {
	Path string
	Name string
}

func (p *Project) Exists() bool {
	projPath := path.Join(os.Getenv("GOPATH"), "src", p.Name)
	_, err := os.Stat(projPath)

	if err == nil {
		return true
	}

	p.Path = projPath
	return false
}

func NewProject(name string) *Project {
	proj := &Project{Name: name}

	if proj.Exists() {
		panic("Project already exists")
	}

	os.Mkdir(proj.Path, 0755)
	os.Mkdir(path.Join(proj.Path, "app"), 0755)
	os.Mkdir(path.Join(proj.Path, "temp"), 0755)

	// dentro da pasta App
	os.Mkdir(path.Join(proj.Path, "app", "controllers"), 0755)
	os.Mkdir(path.Join(proj.Path, "app", "database"), 0755)
	os.Mkdir(path.Join(proj.Path, "app", "middlewares"), 0755)
	os.Mkdir(path.Join(proj.Path, "app", "models"), 0755)
	os.Mkdir(path.Join(proj.Path, "app", "routers"), 0755)

	file, _ := os.Create(path.Join(proj.Path, "app", "main.go"))

	file.WriteString(`package main

import (
	"` + proj.Name + `/app/database"
	// "` + proj.Name + `/app/models"
	"` + proj.Name + `/app/routers"
)

func main() {
	database.StartDB()
	// models.CreateUserTable()
	routers.Router.Run(":8090")
}
`)

	file, _ = os.Create(path.Join(proj.Path, "app", "controllers", "controller.go"))

	file.WriteString(`package controllers

import "github.com/gin-gonic/gin"

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
`)

	file, _ = os.Create(path.Join(proj.Path, "app", "database", "conn.go"))

	file.WriteString(`package database

import (
	"database/sql"
	"log"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

func initDb() *gorp.DbMap {
	db, err := sql.Open("mysql", "user:pass@tcp(localhost:3306)/database")

	if err != nil {
		log.Fatalln(err, "sql.Open failed")
	}

	dbmap := &gorp.DbMap{
		Db: db,
		Dialect: gorp.MySQLDialect{
			Engine:   "InnoDB",
			Encoding: "UTF8",
		},
	}

	return dbmap
}

func StartDB() {
	DBMap = initDb()
}

var DBMap *gorp.DbMap
	
`)

	file, _ = os.Create(path.Join(proj.Path, "app", "routers", "api.go"))

	file.WriteString(`package routers

import (
	"` + proj.Name + `/app/controllers"
	// "` + proj.Name + `/app/middlewares"

	"github.com/gin-gonic/gin"
)

func init() {
	createRouting()
}

var Router *gin.Engine

func createRouting() {
	Router = gin.Default()
	Router.Use(controllers.Cors())
	// Router.Use(gin.Recovery())
	
	/*v1 := Router.Group("/v1")
	{
		v1.GET("/users/", controllers.ListUsers)
		v1.POST("/users/", middlewares.TokenMiddleware(), controllers.AddUser)
		v1.GET("/users/:id", controllers.GetUser)
		v1.PUT("/users/:id", middlewares.TokenMiddleware(), controllers.UpdateUser)
		v1.DELETE("/users/:id", middlewares.TokenMiddleware(), controllers.DeleteUser)
	}*/

}
		
`)

	file.Close()
	return proj
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

func Creator(args []string) {
	switch strings.ToLower(args[0]) {
	case "project":
		NewProject(args[1])
	case "middleware":
		NewMiddleware(args[1])
	case "controller":
		NewController(args[1])
	default:
		panic("Invalid property to create: " + args[1])
	}
}
