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

const text string = "Hello :D"

func init() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, text)
	})
	log.Println("about to list and serve...")
	go func() {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()
}

//
// Testing
//

func Test_GetURLContents(t *testing.T) {
	url := "http://localhost:8080/"
	log.Println("starting the test...")

	contents, err := GetURLContents(url)
	if err != nil {
		t.Errorf("Error from GetURLContents('%s'): %s", url, err)
	}

	response := []byte(text + "\n")

	if string(contents) != string(response) {
		t.Errorf("Got '%q' from GetURLContents('%q'), should be: %q", contents, url, response)
	}
}
