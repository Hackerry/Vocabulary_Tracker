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