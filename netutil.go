package netutil

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetURLContents(url string) (contents []byte, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	contents, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return
}

func RemoteUnmarshal(url string, response interface{}) error {
	contents, err := GetURLContents(url)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(contents, &response); err != nil {
		return err
	}
	return nil
}
