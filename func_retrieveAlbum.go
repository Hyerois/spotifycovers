package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func retrieveAlbum(albumId string, token Token) (SimplifiedAlbum, error) {
	var album SimplifiedAlbum
	url := "https://api.spotify.com/v1/albums/" + albumId
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return album, err
	}
	req.Header.Add("Authorization", "Bearer "+token.String)

	res, err := client.Do(req)
	if err != nil {
		return album, err
	}

	if res.StatusCode != 200 {
		return album, fmt.Errorf("404")
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return album, err
	}

	json.Unmarshal([]byte(body), &album)
	return album, nil
}
