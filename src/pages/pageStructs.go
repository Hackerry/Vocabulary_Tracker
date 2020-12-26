package pages

import (
	"log"
	"path/filepath"
	"html/template"
)

const PageDirectory = "pages"

type WelcomePage struct {
	Title string
	Text string
}

func GetTemplate(fileName string) *template.Template {
	temp, err := template.ParseFiles(filepath.Join(PageDirectory, fileName))
	if err != nil {
		log.Println("Can't get template:", fileName)
	}

	return temp
}