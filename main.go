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
	p := strings.Split(string(data), " ")
	a := api.NewAPI(p[0], p[1])
	log.Println("API URL:", a.Url)
	log.Println("Using parser:", p[1])

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