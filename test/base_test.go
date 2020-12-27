package test

import (
	"time"
	"net/http"
	"strconv"
	"testing"

	"github.com/Hackerry/Vocabulary_Tracker/server"
)

func TestHomePage(t *testing.T) {
	t.Logf("Test server root accessible\n")

	s := server.NewServer(nil)
	go s.Serve()
	time.Sleep(1 * time.Second)

	// Connect to port and 
	resp, err := http.Get("http://localhost:" + strconv.Itoa(server.Port))
	if err != nil {
		t.Errorf("GET failed: %s\n", err)
		t.Fail()
	}
	if resp.StatusCode == 200 {
		t.Logf("GET successful\n")
	}
	resp.Body.Close()
}