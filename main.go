package main

import (
	"log"
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
	a := api.NewAPI(string(data))
	log.Println("API URL:", a.Url)

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