package creator

import (
	"os"
	"path"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Model struct {
	Path string
	Name string
}

func (m *Model) Exists() bool {
	folderPath, _ := os.Getwd()
	modelPath := path.Join(folderPath, "models")

	_, err := os.Stat(modelPath)

	if err != nil {
		return false
	}

	m.Path = modelPath
	return true
}

func NewModel(name string) *Model {
	mod := &Model{Name: name}

	if !mod.Exists() {
		panic("Models folder not found!")
	}

	caser := cases.Title(language.English)

	// dir, _ := os.Getwd()
	// projName := filepath.Base(filepath.Dir(dir))

	file, _ := os.Create(path.Join("models", name+".go"))
	file.WriteString(`package models

import (
	"errors"
	"github.com/local/database"
	"log"
)

type ` + caser.String(name) + ` struct {
	ID        int    ` + "`" + `db:"id, primarykey, autoincrement" json:"id"` + "`" + `
}

func init() {
	Create` + caser.String(name) + `Table()
}

func Create` + caser.String(name) + `Table() {
	database.DBMap.AddTableWithName(` + caser.String(name) + `{}, "` + name + `s")

	err := database.DBMap.CreateTablesIfNotExists()

	database.DBMap.CreateIndex()

	if err != nil {
		log.Fatalln(err, "Could not create ` + name + `s table")
	}
}

// Saves a new ` + name + `
func (` + name + ` *` + caser.String(name) + `) Save() error {
	if database.DBMap.Insert(` + name + `) != nil {
		return errors.New("an error has ocurred on saving a new ` + name + `")
	}

	return nil
}

// Lists all the ` + name + `
func (` + name + ` ` + caser.String(name) + `) List() ([]interface{}, error) {
	return database.DBMap.Select(` + name + `, "select * from ` + name + `s")
}

// Get ` + name + ` by ID
func (` + name + ` *` + caser.String(name) + `) GetById() (interface{}, error) {
	err := database.DBMap.SelectOne(&` + name + `, "select * from ` + name + `s where id = ?", ` + name + `.ID)
	return ` + name + `, err
}

// Checks if ` + name + ` exists
func (` + name + ` *` + caser.String(name) + `) Exists() bool {
	if _, err := ` + name + `.GetById(); err != nil {
		return false
	}

	return true

}

func (` + name + ` *` + caser.String(name) + `) Delete() error {
	if !` + name + `.Exists() {
		return errors.New("` + name + ` not found")
	}

	if response, err := database.DBMap.Exec("DELETE FROM ` + name + `s WHERE id = ?", ` + name + `.ID); err != nil || response == nil {
		return errors.New("could not delete ` + name + `")
	}

	return nil
}
	
`)

	return mod
}
