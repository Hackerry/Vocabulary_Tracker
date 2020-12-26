package server

import (
	"os"
	"log"
	"strconv"
	"net/http"

	"pages"
)

const Port = 8080

func getHomePage(w http.ResponseWriter, req *http.Request) {
	// Go up one level to access data
	os.Chdir("..")

	temp := pages.GetTemplate("home.html")
	data := pages.WelcomePage {
		"Welcome",
		"Enter content here",
	}
	temp.Execute(w, data)
}

func Serve() {
	http.HandleFunc("/", getHomePage)
	log.Fatalf("%s", http.ListenAndServe("localhost:" + strconv.Itoa(Port), nil))
}