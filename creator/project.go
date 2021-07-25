package creator

import (
	"os"
	"path"
)

type Project struct {
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

	envFile, _ := os.Create(path.Join(proj.Path, "app", ".env"))
	file, _ := os.Create(path.Join(proj.Path, "app", "main.go"))

	envFile.WriteString(`
# .env file

PROJECT_NAME=` + name + `
DATABASE=mysql
DB_HOST=localhost
DB_PORT=27017
DB_USERNAME=admin
DB_PASSWORD=password
DB_NAME=testdb
`)

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
	"github.com/joho/godotenv"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

func initDb() *gorp.DbMap {
	var projEnv map[string]string
	projEnv, err := godotenv.Read()

	if err != nil {
		panic(".env file does not exist!")
	}
	
	dbString := projEnv["DB_USERNAME"] + ":" + projEnv["DB_PASSWORD"] + "@tcp(" + projEnv["DB_HOST"] + ":" + projEnv["DB_PORT"] + ")/" + projEnv["DB_NAME"]
	db, err := sql.Open(projEnv["DATABASE"], dbString)

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
