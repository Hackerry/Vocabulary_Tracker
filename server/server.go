package server

import (
	"os"
	"log"
	"time"
	"regexp"
	"strings"
	"strconv"
	"net/http"
	"io/ioutil"
	"path/filepath"

	"github.com/Hackerry/Vocabulary_Tracker/pages"
	"github.com/Hackerry/Vocabulary_Tracker/api"
	"github.com/Hackerry/Vocabulary_Tracker/entryStore"
)

const Port = 8080
const TimeFormat = "2006/01/02 15:04:05"
const CacheDirectory = "cache"
const CachedQueryFile = "_cache.json"

var reg = regexp.MustCompile("^[^\\%?&#!/:.\"=^$]*$")

// TODO cache all page templates before server starting

type Server struct {
	api *api.API
}

// Get home page
func (s *Server) getHomePage(w http.ResponseWriter, req *http.Request) {
	temp := pages.GetTemplate("home.html")
	if temp == nil {
		// TODO send 500
		w.WriteHeader(500)
		w.Write([]byte("500 Error: can't get home page template"))
		return
	}

	// Post request to add new tag
	var data pages.WelcomePage
	if req.Method == "POST" {
		// Create tag
		if err := req.ParseForm(); err != nil {
			log.Println("Failed to parse form in getHomePage", err)
			return
		}
		tag := req.FormValue("addTag")
		bgColor := req.FormValue("bgColor")
		txColor := req.FormValue("txColor")

		// Error creating new tag, report error to user
		_, msg := entryStore.CreateTag(tag, bgColor, txColor)
		data.Message = msg
	}

	// Get all tags if any
	tags := entryStore.ReadTags()
	if tags == nil {
		tags = make([]entryStore.Tag, 0, 0)
	}

	data.Tags = tags
	temp.Execute(w, data)
}

// Get search word page
func (s *Server) getWordPage(w http.ResponseWriter, req *http.Request) {
	// Get word from parameter
	params := req.URL.Query()
	word := params["word"]

	// Sanitize word input to disallow [^\%?&#!/:."=] special characters
	if word == nil || len(word) == 0 || !reg.Match([]byte(word[0])) {
		// TODO send 400
		w.WriteHeader(400)
		w.Write([]byte("400 Error: wrong format"))
		return
	}
	queryWord := strings.TrimSpace(word[0])

	// Get optional tag
	tag := params["tag"]
	var tagToUse *entryStore.Tag
	if tag == nil || len(tag) == 0 {
		// No tag provided
		tagToUse = nil
	} else if readTag := entryStore.ReadTag(tag[0]); readTag != nil {
		// Construct tag
		tagToUse = readTag
	}

	// Post new comment on this word
	if req.Method == "POST" {
		if err := req.ParseForm(); err != nil {
			log.Println("Failed to parse form in getWordPage")
			return
		}
		
		// Create new comment
		comment := req.FormValue("comment")
		t := time.Now()
		date := t.Format(TimeFormat)
		var tagName string
		if tagToUse == nil {
			tagName = ""
		} else {
			tagName = tagToUse.Name
		}
		entry := entryStore.Entry {
			Date:		date,
			Comment:	comment,
			Tag:		tagName,
		}

		// Write entry
		entries := entryStore.ReadEntries(queryWord)
		if entries == nil {
			log.Println("Error reading entry file")
		} else {
			entries = append(entries, entry)
			entryStore.WriteEntries(entries, queryWord)
		}

		// Store word to tag
		if tagName != "" {
			entryStore.AddWord(tagName, queryWord)
		}
	}

	// Search through cache first
	var queryResp []byte
	cacheFileName := queryWord + CachedQueryFile
	if _, err := os.Stat(cacheFileName); os.IsNotExist(err) {
		// Query words
		queryResp = s.api.Query(queryWord)
		if queryResp == nil {
			// TODO send 500
			w.WriteHeader(500)
			w.Write([]byte("500 Error: fail to query for word: " + queryWord))
			return
		}

		// Create cache folder
		if stat, err := os.Stat(CacheDirectory); os.IsNotExist(err) {
			// Directory not exist, create it
			os.Mkdir(CacheDirectory, 0755)
		} else if !stat.Mode().IsDir() {
			// Not directory
			log.Fatal("Error: found \"cache\" file")
		}

		// To save some api usage, cache all words queried
		err := ioutil.WriteFile(filepath.Join(CacheDirectory, cacheFileName), queryResp, 0644)
		if err != nil {
			log.Println("Failed to cache response")
		}
	} else {
		queryResp, err = ioutil.ReadFile(cacheFileName)
		if err != nil {
			// TODO send 500
			w.WriteHeader(500)
			w.Write([]byte("500 Error: fail to read cache file: " + cacheFileName))
			return
		}
	}

	defs := s.api.Parser(queryResp)

	// Uncomment to filter out idioms, prefix & suffixs and only keep words matching exactly as input
	// if defs.DefList != nil {
	// 	queryWordDefs := make([]api.Definition, 0, 0)
	// 	for _, def := range defs.DefList {
	// 		if strings.ToLower(def.Word) == strings.ToLower(queryWord) {
	// 			queryWordDefs = append(queryWordDefs, def)
	// 		}
	// 	}

	// 	defs.DefList = queryWordDefs
	// }

	// Get user entries
	entries := entryStore.ReadEntries(queryWord)
	if entries == nil || len(entries) == 0 {
		entries = make([]entryStore.Entry, 1, 1)
		entries[0] = entryStore.Entry {
			Date: 		"",
			Comment: 	"No Entry Found",
			Tag:		"",
		}
	}

	data := pages.WordPage {
		HaveTag:		tagToUse != nil,
		Tag:			tagToUse,
		Definitions:	defs,
		Entries:		entries,
	}
	
	temp := pages.GetTemplate("search.html")
	if temp == nil {
		// TODO send 500
		w.WriteHeader(500)
		w.Write([]byte("500 Error: can't get search page template"))
		return
	}

	temp.Execute(w, data)
}

// Helper function to get all tags and words
func (s *Server) getTagPage(w http.ResponseWriter, req *http.Request) {
	// Get all tags
	tags := entryStore.ReadTags()
	if tags == nil {
		tags = make([]entryStore.Tag, 0, 0)
	}

	// Get words for the selected tag
	params := req.URL.Query()
	tag := params["tag"]

	var words []string
	tagName := ""
	if tag == nil || len(tag) == 0 {
		words = make([]string, 0, 0)
	} else {
		// Get words from the current tag
		tagName = strings.TrimSpace(tag[0])
		for _, t := range tags {
			if t.Name == tagName {
				words = t.Words
			}
		}
	}

	data := pages.TagPage {
		Tags: tags,
		Words: words,
		CurrTag: tagName,
	}

	temp := pages.GetTemplate("tags.html")
	if temp == nil {
		// TODO send 500
		w.WriteHeader(500)
		w.Write([]byte("500 Error: can't get tag page template"))
		return
	}
	temp.Execute(w, data)
}

func (s *Server) Serve() {
	// Register all operations
	http.HandleFunc("/", s.getHomePage)
	http.HandleFunc("/search", s.getWordPage)
	http.HandleFunc("/tag", s.getTagPage)

	// Clear cache on start
	_ = os.RemoveAll(CacheDirectory)
	
	// Starting server
	log.Println("Starting server...")
	log.Fatalf("%s", http.ListenAndServe("localhost:" + strconv.Itoa(Port), nil))
}

func NewServer(api *api.API) *Server {
	return &Server{ api };
}