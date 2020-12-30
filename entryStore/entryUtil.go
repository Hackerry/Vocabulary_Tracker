package entryStore

import (
	"os"
	"log"
	"strings"
	"io/ioutil"
	"encoding/json"
	"path/filepath"
)

const EntryDirectory = "entries"

type Entry struct {
	Date	string
	Comment	string
	Tag		string
}

// Helper method that reads entries
func ReadEntries(word string) []Entry {
	// Check folder exists
	if stat, err := os.Stat(EntryDirectory); os.IsNotExist(err) {
		// Directory not exist, create it
		os.Mkdir(EntryDirectory, 0755)
	} else if !stat.Mode().IsDir() {
		// Not directory
		log.Fatal("Error: found \"entries\" file")
	}

	word = strings.TrimSpace(word)
	if word == "" || strings.Index(word, "/") != -1 || strings. Index(word, "\\") != -1 {
		return nil
	}

	// Retrieve entries
	var entries []Entry
	fileName := filepath.Join(EntryDirectory, word) + ".json"

	// Possible file doesn't exist, return an empty list
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return make([]Entry, 0, 0)
	}

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println("Can't read entry file:", fileName, "-", err)
		return nil
	}
	err = json.Unmarshal(data, &entries)
	if err != nil {
		log.Println("Error parsing json file:", fileName, "-", err)
		return nil
	}

	return entries
}

// Helper method that writes entries
func WriteEntries(entries []Entry, word string) bool {
	word = strings.TrimSpace(word)
	if word == "" || strings.Index(word, "/") != -1 || strings. Index(word, "\\") != -1 {
		return false
	}

	// Encode into JSON
	data, err := json.Marshal(entries)
	if err != nil {
		log.Println("Error marshaling entries to JSON:", word)
		return false
	}

	// Write to file
	fileName := filepath.Join(EntryDirectory, word) + ".json"
	err = ioutil.WriteFile(fileName, data, 0755)
	if err != nil {
		log.Println("Fail to write to file", fileName)
		return false
	}

	return true
}