package entryStore

import (
	"os"
	"log"
	"regexp"
	"strings"
	"io/ioutil"
	"encoding/json"
	"path/filepath"
)

const TagDirectory = "tags"
const JSONExt = ".json"

type Tag struct {
	Name	string
	BgColor	string
	TxColor	string
	Words	[]string
}

var colorReg = regexp.MustCompile("^#[0123456789abcdef]{6}$")

// Read all existing tag names
func ReadTags() []Tag {
	// Check folder exists
	if stat, err := os.Stat(TagDirectory); os.IsNotExist(err) {
		// Directory not exist, create it
		os.Mkdir(TagDirectory, 0755)
		return nil
	} else if !stat.Mode().IsDir() {
		// Not directory
		log.Fatal("Error: found \"tags\" file")
	}

	listTags := make([]Tag, 0, 0)
	err := filepath.Walk(TagDirectory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			// Skip directories
			return nil
		}

		// Must be a json file
		if filepath.Ext(info.Name()) == JSONExt {
			fileName := filepath.Base(info.Name())
			tagName := fileName[:len(fileName)-len(JSONExt)]
			tag := ReadTag(tagName)
			if tag == nil {
				log.Println("Can't read tag file:", tag)
			} else {
				listTags = append(listTags, *tag)
			}
		}
		
		return nil
	})
	if err != nil {
		log.Println("Error retrieving tags", err)
		return nil
	}

	return listTags
}

// Helper function to read a tag file
func ReadTag(tag string) *Tag {
	tag = strings.TrimSpace(tag)
	if tag == "" {
		return nil
	}

	fileName := filepath.Join(TagDirectory, tag + JSONExt)

	// Can't find tag file
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		log.Println("Can't read tag file:", tag, "-", err)
		return nil
	}

	// Retrieve tag
	var readTag Tag
	data, err := ioutil.ReadFile(fileName)
	err = json.Unmarshal(data, &readTag)
	if err != nil {
		log.Println("Error parsing json file:", fileName, "-", err)
		return nil
	}

	return &readTag
}

// Helper function to write a tag file
func AddWord(tag string, word string) {
	readTag := ReadTag(tag)
	if readTag == nil {
		log.Println("Can't get tag file:", tag)
		return
	}

	for _, w := range readTag.Words {
		// Word already tagged
		if w == word {
			return
		}
	}

	readTag.Words = append(readTag.Words, word)

	// Store word
	fileName := filepath.Join(TagDirectory, tag + ".json")
	data, err := json.Marshal(readTag)
	if err != nil {
		log.Println("Error marshaling json file for tag:", tag)
		return
	}
	ioutil.WriteFile(fileName, data, 0644)
}

// Helper function to create new tags
func CreateTag(tag string, bgColor string, textColor string) (*Tag, string) {
	tag = strings.TrimSpace(tag)
	if tag == "" {
		return nil, "Can't use empty string as tag name."
	} else if strings.Index(tag, "/") != -1 || strings. Index(tag, "\\") != -1 {
		return nil, "Can't use path separator('\\', '/') characters as tag name."
	}

	// Check if tag already exists
	fileName := filepath.Join(TagDirectory, tag + ".json")
	if _, err := os.Stat(fileName); !os.IsNotExist(err) {
		return nil, "Tag with name " + tag + " already exists."
	}

	// Check color format correct: #xxxxxx
	if !colorReg.Match([]byte(bgColor)) || !colorReg.Match([]byte(textColor)) {
		return nil, "Color must be of format: #xxxxxx"
	}

	// Default color is black and white
	bgColor = strings.TrimSpace(bgColor)
	textColor = strings.TrimSpace(textColor)
	if bgColor == "" || textColor == "" {
		bgColor = "#000000"
		textColor = "#ffffff"
	}

	// Create new tag
	newTag := Tag{
		Name: tag,
		BgColor: bgColor,
		TxColor: textColor,
		Words: make([]string, 0, 0),
	}

	// Store tag structure
	data, err := json.Marshal(newTag)
	if err != nil {
		return nil, "Error marshaling json file for tag:" + tag
	}
	ioutil.WriteFile(fileName, data, 0644)

	return &newTag, ""
}