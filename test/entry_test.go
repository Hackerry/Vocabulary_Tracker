package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Hackerry/Vocabulary_Tracker/entryStore"
)

func TestReadEntry(t *testing.T) {
	t.Logf("Test read entry file\n")

	entries := entryStore.ReadEntries("some")

	if entries == nil {
		t.Errorf("Should return entries")
	}
	if len(entries) != 3 {
		t.Errorf("Should return 3 entries")
	}
	if entries[0].Date != "01/20/2019" || entries[0].Comment != "saw in NYTimes" {
		t.Errorf("Entry 1 mismatch %v\n", entries[0])
	}
	if entries[1].Date != "02/01/2020" || entries[1].Comment != "i should have remembered it" {
		t.Errorf("Entry 2 mismatch %v\n", entries[1])
	}
	if entries[2].Date != "02/20/2020" || entries[2].Comment != "again" {
		t.Errorf("Entry 3 mismatch %v\n", entries[2])
	}

	entries = entryStore.ReadEntries("non-exist")
	if entries != nil {
		t.Errorf("Should return nil")
	}

	entries = entryStore.ReadEntries("")
	if entries != nil {
		t.Errorf("Should return nil")
	}
}

func TestWriteEntry(t *testing.T) {
	t.Logf("Test write entry file\n")

	word := "comment"
	fileName := filepath.Join(entryStore.EntryDirectory, word) + ".json"
	_ = os.Remove(fileName)

	entries := []entryStore.Entry {
		entryStore.Entry{"09/01/2010", "comment 0"},
		entryStore.Entry{"10/01/2010", "comment 1"},
		entryStore.Entry{"11/01/2010", "comment 2"},
	}

	entryStore.WriteEntries(entries, word)

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		t.Fatalf("Should create entry file")
	}

	entries = entryStore.ReadEntries(word)
	if entries == nil {
		t.Errorf("Should return entries")
	}

	for idx, e := range entries {
		if e != entries[idx] {
			t.Errorf("Entries[%v] is not equal: %v\n", idx, entries[idx])
		}
	}
}

func TestSpecialPaths(t *testing.T) {
	t.Logf("Test special entry names")

	entries := entryStore.ReadEntries("")
	if entries != nil {
		t.Errorf("Should be nil")
	}
	entries = entryStore.ReadEntries("../something")
	if entries != nil {
		t.Errorf("Should be nil")
	}
	entries = entryStore.ReadEntries("C:\\something")
	if entries != nil {
		t.Errorf("Should be nil")
	}
	entries = entryStore.ReadEntries("/bin/bash")
	if entries != nil {
		t.Errorf("Should be nil")
	}
}