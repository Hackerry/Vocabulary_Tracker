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
	Tags	[]entryStore.Tag
	Message	string
}

type WordPage struct {
	HaveTag		bool
	Tag			*entryStore.Tag
	Definitions	*api.Definitions
	Entries		[]entryStore.Entry
}

type TagPage struct {
	Tags	[]entryStore.Tag
	CurrTag	string
	Words	[]string
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