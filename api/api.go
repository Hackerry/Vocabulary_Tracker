package api

import (
	"log"
	"strings"
	"net/url"
	"net/http"
	"io/ioutil"
)

const WordPlaceHolder = "[word]"

type API struct {
	Url			string
	Parser		Parser
}

func (api *API) Query(word string) []byte {
	query := strings.ReplaceAll(api.Url, WordPlaceHolder, word)
	log.Println("Query:", query)
	resp, err := http.Get(query)
	if err != nil {
		log.Println("Error calling API", err)
		return nil
	}
	defer resp.Body.Close()

	// Read in body
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading API response", err)
		return nil
	}

	// Check response code
	if resp.StatusCode != 200 {
		log.Println("Response not 200:")
		log.Println(string(data))
	}

	return data
}

// Return new API with the format: http://www.xxxx.com?word=[word]&key=...
func NewAPI(urlStr string, apiProvider string) *API {
	u, err := url.Parse(urlStr)
	if err != nil {
		log.Println("Error parsing input URL", err)
		return nil
	}

	// Force scheme to be HTTP
	u.Scheme = "http"

	// Check "[word]" place holder
	if !strings.Contains(u.String(), WordPlaceHolder) {
		log.Println("Error: No word query string", u.Path)
		return nil
	}

	// Link parser to apiProvider
	var parser Parser
	switch apiProvider {
		case "Merriam-Webster":
			parser = ParseMWJson
		default:
			log.Println("No parser found")
			return nil
	}

	return &API {
		Url:	u.String(),
		Parser:	parser,
	};
}