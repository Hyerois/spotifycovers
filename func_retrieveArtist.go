package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func retrieveArtist(artistId string, token Token) (Artist, error) {
	var artist Artist
	url := "https://api.spotify.com/v1/artists/" + artistId
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return artist, err
	}
	req.Header.Add("Authorization", "Bearer "+token.String)

	res, err := client.Do(req)
	if err != nil {
		return artist, err
	}

	if res.StatusCode != 200 {
		return artist, fmt.Errorf("404")
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return artist, err
	}

	json.Unmarshal([]byte(body), &artist)
	return artist, nil
}
