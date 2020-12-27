package server

import (
	"log"
	"strconv"
	"net/http"

	"github.com/Hackerry/Vocabulary_Tracker/pages"
	"github.com/Hackerry/Vocabulary_Tracker/api"
)

const Port = 8080

type Server struct {
	api *api.API
}

func (s *Server) getHomePage(w http.ResponseWriter, req *http.Request) {
	temp := pages.GetTemplate("home.html")
	if temp == nil {
		w.WriteHeader(500)
		w.Write([]byte("500 Error"))
		return
	}

	data := pages.WelcomePage {
		"Welcome",
		"Enter content here",
	}
	temp.Execute(w, data)
}

func (s *Server) Serve() {
	// Register all operations
	http.HandleFunc("/", s.getHomePage)
	
	// Starting server
	log.Println("Starting server...")
	log.Fatalf("%s", http.ListenAndServe("localhost:" + strconv.Itoa(Port), nil))
}

func NewServer(api *api.API) *Server {
	return &Server{ api };
}