package netutil

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

//
// Test setup
//

// replace this with net/http/httptest

const text string = "Hello :D"

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprint(w, text)
	})
	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
}

//
// Testing
//

func TestGetURLContents(t *testing.T) {
	url := "http://localhost:8080/"

	contents, err := GetURLContents(url)
	if err != nil {
		t.Fatalf("Error from GetURLContents('%q'): %q", url, err)
	}

	response := []byte(text)

	if string(contents) != string(response) {
		t.Errorf("Got '%q' from GetURLContents('%q'), should be: %q", contents, url, response)
	}
}

func TestRemoteUnmarshal(t *testing.T) {
	// TODO
}
