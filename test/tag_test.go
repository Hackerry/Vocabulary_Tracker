package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Hackerry/Vocabulary_Tracker/entryStore"
)

func TestReadTagNames(t *testing.T) {
	t.Logf("Test read tag files\n")

	tags := entryStore.ReadTags()

	if tags == nil || len(tags) != 2 {
		t.Errorf("Failed to retrieve tags %v\n", tags)
	}

	if tags[0] != "NYTimes" && tags[1] != "NYTimes" {
		t.Errorf("No NYTimes tag")
	}
	if tags[0] != "Sherlock Holmes" && tags[1] != "Sherlock Holmes" {
		t.Errorf("No Sherlock Holmes tag")
	}
}

func TestReadTag(t *testing.T) {
	t.Logf("Test read individual files")

	// Read tag
	tag := entryStore.ReadTag("Sherlock Holmes")
	temp := entryStore.Tag {
		Name: "Sherlock Holmes",
		BgColor: "#000000",
		TxColor: "#ffffff",
		Words: []string {
			"page", "wagon", "speckled",
		},
	}
	if tag == nil || tag.Name != temp.Name || tag.BgColor != temp.BgColor || tag.TxColor != temp.TxColor {
		t.Errorf("Retrieved tag is different: %v - %v\n", tag, temp)
	}
	for _, s := range tag.Words {
		matched := false
		for _, ss := range temp.Words {
			if s == ss {
				matched = true
				break
			}
		}

		if !matched {
			t.Errorf("Word %s mismatch", s)
		}
	}

	if len(tag.Words) != len(temp.Words) {
		t.Errorf("Word length mismatch")
	}

	tag = entryStore.ReadTag("NYTimes")
	temp = entryStore.Tag {
		Name: "NYTimes",
		BgColor: "#000000",
		TxColor: "#ffffff",
		Words: []string {
			"hegemony", "propaganda",
		},
	}
	if tag == nil || tag.Name != temp.Name || tag.BgColor != temp.BgColor || tag.TxColor != temp.TxColor {
		t.Errorf("Retrieved tag is different: %v - %v\n", tag, temp)
	}
	for _, s := range tag.Words {
		matched := false
		for _, ss := range temp.Words {
			if s == ss {
				matched = true
				break
			}
		}

		if !matched {
			t.Errorf("Word %s mismatch\n", s)
		}
	}

	if len(tag.Words) != len(temp.Words) {
		t.Errorf("Word length mismatch\n")
	}

	tag = entryStore.ReadTag("something")
	if tag != nil {
		t.Errorf("Should return nil\n")
	}
}

func TestCreateTag(t *testing.T) {
	fileName := filepath.Join(entryStore.TagDirectory, "newTag" + entryStore.JSONExt)

	newTag := entryStore.Tag {
		Name: "newTag",
		BgColor: "#123455",
		TxColor: "#232323",
		Words: []string {
			"word1", "word2",
		},
	}
	tag, msg := entryStore.CreateTag(newTag.Name, newTag.BgColor, newTag.TxColor)
	if tag == nil {
		t.Errorf("Error creating new tag: %v with message %v\n", tag, msg)
		return
	}

	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		t.Errorf("Can't find created tag file\n")
		return
	}

	readTag := entryStore.ReadTag("newTag")
	if readTag == nil {
		t.Errorf("Can't read created tag file\n")
		return
	}

	if readTag.Name != newTag.Name || readTag.BgColor != newTag.BgColor || readTag.TxColor != newTag.TxColor {
		t.Errorf("Retrieved tag is different: %v - %v\n", readTag, newTag)
	}
	for _, s := range readTag.Words {
		matched := false
		for _, ss := range newTag.Words {
			if s == ss {
				matched = true
				break
			}
		}

		if !matched {
			t.Errorf("Word %s mismatch\n", s)
		}
	}

	if readTag.Words == nil || len(readTag.Words) != 0 {
		t.Errorf("Word length should be 0\n")
	}

	_ = os.Remove(fileName)

	// Create existing tag
	tag, msg = entryStore.CreateTag("NYTimes", "", "")
	if tag != nil {
		t.Errorf("Should be nil")
	}

	// Create special tag
	tag, msg = entryStore.CreateTag("", "", "")
	if tag != nil {
		t.Errorf("Should be nil")
	}
	tag, msg = entryStore.CreateTag("../hello", "", "")
	if tag != nil {
		t.Errorf("Should be nil")
	}
	tag, msg = entryStore.CreateTag("C:\\hello\\hi", "", "")
	if tag != nil {
		t.Errorf("Should be nil")
	}
}

func TestAddWord(t *testing.T) {
	fileName := filepath.Join(entryStore.TagDirectory, "newTag" + entryStore.JSONExt)

	newTag := entryStore.Tag {
		Name: "newTag",
		BgColor: "#123455",
		TxColor: "#232323",
		Words: []string {
			"word1", "word2",
		},
	}
	tag, _ := entryStore.CreateTag(newTag.Name, newTag.BgColor, newTag.TxColor)
	if tag == nil {
		t.Errorf("Tag can't be created")
		return
	}
	
	// Add tags
	for _, w := range newTag.Words {
		entryStore.AddWord("newTag", w)
	}

	readTag := entryStore.ReadTag("newTag")
	if readTag == nil {
		t.Errorf("Can't read created tag file\n")
		return
	}

	if readTag.Name != newTag.Name || readTag.BgColor != newTag.BgColor || readTag.TxColor != newTag.TxColor {
		t.Errorf("Retrieved tag is different: %v - %v\n", readTag, newTag)
	}
	for _, s := range readTag.Words {
		matched := false
		for _, ss := range newTag.Words {
			if s == ss {
				matched = true
				break
			}
		}

		if !matched {
			t.Errorf("Word %s mismatch\n", s)
		}
	}

	if readTag.Words == nil || len(readTag.Words) != len(newTag.Words) {
		t.Errorf("Word length mismatch: %v - %v\n", readTag.Words, newTag.Words)
	}

	_ = os.Remove(fileName)
}