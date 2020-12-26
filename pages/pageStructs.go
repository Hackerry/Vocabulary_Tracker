package pages

import (
	"log"
	"path/filepath"
	"html/template"
)

var PageDirectory, _ = filepath.Rel(".", "Vocabulary_Tracker/pages")

type WelcomePage struct {
	Title string
	Text string
}

func GetTemplate(fileName string) *template.Template {
	log.Println("Get template", fileName)
	f := filepath.Join(PageDirectory, fileName)
	temp, err := template.ParseFiles(f)
	if err != nil {
		log.Println("Can't parse template:", fileName, err)
		return nil
	}

	return temp
}