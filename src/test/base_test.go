package test

import (
	"time"
	"net/http"
	"strconv"
	"testing"

	"server"
)

func TestBase(t *testing.T) {
	t.Logf("Base test")
}

func TestHomePage(t *testing.T) {
	go server.Serve()
	time.Sleep(1 * time.Second)

	// Connect to port and 
	resp, err := http.Get("http://localhost:" + strconv.Itoa(server.Port))
	if err != nil {
		t.Errorf("GET failed: %s\n", err)
	}
	if resp.StatusCode == 200 {
		t.Logf("GET successful\n")
	}
}

func TestHomePage(t *testing.T) {
	go server.Serve()
	time.Sleep(1 * time.Second)

	// Connect to port and 
	resp, err := http.Get("http://localhost:" + strconv.Itoa(server.Port))
	if err != nil {
		t.Errorf("GET failed: %s\n", err)
	}
	if resp.StatusCode == 200 {
		t.Logf("GET successful\n")
	}
}