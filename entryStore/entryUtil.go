package entryStore

import (
	"os"
	"log"
	"io/ioutil"
	"encoding/json"
	"path/filepath"
)

const EntryDirectory = "entries"

type Entry struct {
	Date	string
	Comment	string
}

func ReadEntries(word string) []Entry {
	// Check folder exists
	if stat, err := os.Stat(EntryDirectory); os.IsNotExist(err) {
		// Directory not exist, create it
		os.Mkdir(EntryDirectory, 0755)
	} else if !stat.Mode().IsDir() {
		// Not directory
		log.Fatal("Error: found \"entries\" file")
	}

	if word == "" {
		return nil
	}

	// Retrieve entries
	var entries []Entry
	fileName := filepath.Join(EntryDirectory, word) + ".json"
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

func WriteEntries(entries []Entry, word string) {
	if word == "" {
		return
	}

	// Encode into JSON
	data, err := json.Marshal(entries)
	if err != nil {
		log.Println("Error marshaling entries to JSON:", word)
		return
	}

	// Write to file
	fileName := filepath.Join(EntryDirectory, word) + ".json"
	ioutil.WriteFile(fileName, data, 0755)
}