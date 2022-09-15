package creator

import (
	"os"
	"os/exec"
	"path"
	"runtime"
)

type Project struct {
	Path string
	Name string
}

func (p *Project) Exists() bool {
	projPath := path.Join(p.Name)
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

	createFolders(proj.Path)

	createEnvFile(proj.Path, proj.Name)
	createMainFile(proj.Path)
	createControllerCorsFile(proj.Path)

	// To Run Dockerfile: docker build -t abc -f Dockerfile .
	createDockerfile(proj.Path)
	createRouterFile(proj.Path)
	createDatabaseConnectionFile(proj.Path)
	createAndUpdateModulesFile(proj.Path)

	return proj
}

func createFolders(projectPath string) {
	os.Mkdir(projectPath, 0755)
	os.Mkdir(path.Join(projectPath, "app"), 0755)

	// dentro da pasta App
	os.Mkdir(path.Join(projectPath, "app", "controllers"), 0755)
	os.Mkdir(path.Join(projectPath, "app", "database"), 0755)
	os.Mkdir(path.Join(projectPath, "app", "middlewares"), 0755)
	os.Mkdir(path.Join(projectPath, "app", "models"), 0755)
	os.Mkdir(path.Join(projectPath, "app", "routers"), 0755)
}

func createEnvFile(projectPath string, projectName string) {
	file, _ := os.Create(path.Join(projectPath, "app", ".env"))

	file.WriteString(`
# .env file

PROJECT_NAME=` + projectName + `
DATABASE=mysql
DB_HOST=localhost
DB_PORT=27017
DB_USERNAME=admin
DB_PASSWORD=password
DB_NAME=testdb
`)
	file.Close()
}

func createMainFile(projectPath string) {
	file, _ := os.Create(path.Join(projectPath, "app", "main.go"))
	file.WriteString(`package main

import (
	"github.com/local/routers"
)

func main() {
	routers.Router.Run(":8090")
}
`)
	file.Close()
}

func createControllerCorsFile(projectPath string) {
	file, _ := os.Create(path.Join(projectPath, "app", "controllers", "controller.go"))

	file.WriteString(`package controllers

import "github.com/gin-gonic/gin"

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
`)

	file.Close()
}

func createAndUpdateModulesFile(projectPath string) {
	file, _ := os.Create(path.Join(projectPath, "app", "go.mod"))
	file.WriteString(`module app

go ` + runtime.Version()[2:6] + `


replace github.com/local => ./

	
`)

	file.Close()

	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = path.Join(projectPath, "app")

	if err := cmd.Start(); err != nil {
		panic("Was not possible to start CMD")
	}
}

func createDatabaseConnectionFile(projectPath string) {
	file, _ := os.Create(path.Join(projectPath, "app", "database", "conn.go"))

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

func init() {
	StartDB()
}

var DBMap *gorp.DbMap
	
`)

	file.Close()

}

func createRouterFile(projectPath string) {
	file, _ := os.Create(path.Join(projectPath, "app", "routers", "api.go"))

	file.WriteString(`package routers

import (
	"github.com/local/controllers"
	// "github.com/local//middlewares"

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

}

func createDockerfile(projectPath string) {
	file, _ := os.Create(path.Join(projectPath, "Dockerfile"))

	file.WriteString(`FROM golang:` + runtime.Version()[2:6] + `-alpine

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app/go-sample-app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY app/go.mod .
COPY app/go.sum .

RUN go mod download

COPY app/ .

# Build the Go app
RUN go build -o ./out/go-sample-app .


# This container exposes port 8080 to the outside world
EXPOSE 8090

# Run the binary program produced by 'go install'
CMD ["./out/go-sample-app"]
`)

	file.Close()

}
