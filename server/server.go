package server

import (
	"os"
	"log"
	"strconv"
	"net/http"

	"github.com/Hackerry/Vocabulary_Tracker/pages"
)

const Port = 8080

func getHomePage(w http.ResponseWriter, req *http.Request) {
	// Go up one level to access data
	os.Chdir("..")

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

func Serve() {
	http.HandleFunc("/", getHomePage)
	
	log.Println("Starting server...")
	log.Fatalf("%s", http.ListenAndServe("localhost:" + strconv.Itoa(Port), nil))
}