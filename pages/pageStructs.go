package pages

import (
	"log"
	"path/filepath"
	"html/template"

	"github.com/Hackerry/Vocabulary_Tracker/api"
	"github.com/Hackerry/Vocabulary_Tracker/entryStore"
)

const PageDirectory = "pages"

type WelcomePage struct {
	Title		string
	Text		string
}

type WordPage struct {
	Definitions	*api.Definitions
	Entries		[]entryStore.Entry
}

// Retrieve template file
func GetTemplate(fileName string) *template.Template {
	f := filepath.Join(PageDirectory, fileName)
	temp, err := template.ParseFiles(f)
	if err != nil {
		log.Println("Can't parse template:", fileName, err)
		return nil
	}

	return temp
}