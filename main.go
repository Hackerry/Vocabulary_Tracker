package main

import (
	"log"
	"strings"
	"io/ioutil"

	"github.com/Hackerry/Vocabulary_Tracker/server"
	"github.com/Hackerry/Vocabulary_Tracker/api"
)

func main() {
	// Get api
	data, err := ioutil.ReadFile("api_key")
	if err != nil {
		log.Println(err)
	}
	lines := strings.Split(string(data), "\n")

	var a *api.API = nil
	// Comment out unused API with #
	for _, line := range lines {
		if line[0] != '#' {
			p := strings.Split(line, " ")
			a = api.NewAPI(p[0], p[1])
			log.Println("API URL:", a.Url)
			log.Println("Using parser:", p[1])
			break
		}
	}

	if a == nil {
		log.Fatal("Can't initialize api\n")
	}

	s := server.NewServer(a)
	s.Serve()

	// Test get word
	// resp := a.Query("voluminous")
	// if resp != nil {
	// 	log.Println(string(resp))
	// } else {
	// 	log.Println("Empty")
	// }
}