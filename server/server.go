package server

import (
	"log"
	"regexp"
	"strings"
	"strconv"
	"net/http"

	"github.com/Hackerry/Vocabulary_Tracker/pages"
	"github.com/Hackerry/Vocabulary_Tracker/api"
)

const Port = 8080

var reg = regexp.MustCompile("^[^\\%?&#!/:.'\"=^$]*$")

// TODO cache all page templates before server starting

type Server struct {
	api *api.API
}

func (s *Server) getHomePage(w http.ResponseWriter, req *http.Request) {
	temp := pages.GetTemplate("home.html")
	if temp == nil {
		// TODO send 500
		w.WriteHeader(500)
		w.Write([]byte("500 Error: can't get home page template"))
		return
	}

	data := pages.WelcomePage {
		"Welcome",
		"Enter content here",
	}
	temp.Execute(w, data)
}

func (s *Server) searchWord(w http.ResponseWriter, req *http.Request) {
	// Query for word
	params := req.URL.Query()
	word := params["word"]

	// Sanitize word input to disallow [^\%?&#!/:.'"=] special characters
	if word == nil || len(word) == 0 || !reg.Match([]byte(word[0])) {
		// TODO send 400
		w.WriteHeader(400)
		w.Write([]byte("400 Error: wrong format"))
		return
	}

	// Query words
	queryWord := strings.TrimSpace(word[0])
	data := s.api.Query(queryWord)
	if data == nil {
		// TODO send 500
		w.WriteHeader(500)
		w.Write([]byte("500 Error: fail to query for word"))
		return
	}

	// Parse response, filter out idioms, prefix & suffixs
	defs := s.api.Parser(data)
	if defs.DefList != nil {
		queryWordDefs := make([]api.Definition, 0, 0)
		for _, def := range defs.DefList {
			if strings.ToLower(def.Word) == strings.ToLower(queryWord) {
				queryWordDefs = append(queryWordDefs, def)
			}
		}

		defs.DefList = queryWordDefs
	}
	
	temp := pages.GetTemplate("search.html")
	if temp == nil {
		// TODO send 500
		w.WriteHeader(500)
		w.Write([]byte("500 Error: can't get search page template"))
		return
	}

	temp.Execute(w, defs)
}

func (s *Server) Serve() {
	// Register all operations
	http.HandleFunc("/", s.getHomePage)
	http.HandleFunc("/search", s.searchWord)
	
	// Starting server
	log.Println("Starting server...")
	log.Fatalf("%s", http.ListenAndServe("localhost:" + strconv.Itoa(Port), nil))
}

func NewServer(api *api.API) *Server {
	return &Server{ api };
}