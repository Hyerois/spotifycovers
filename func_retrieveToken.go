package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func retrieveToken() (Token, error) {
	var token Token

	payload := strings.NewReader("grant_type=client_credentials")

	client := &http.Client{}
	req, err := http.NewRequest("POST", baseApi, payload)

	if err != nil {
		return token, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Authorization", "Basic [SECRET]")

	res, err := client.Do(req)
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return token, err
	}
	defer res.Body.Close()

	json.Unmarshal([]byte(body), &token)
	return token, nil
}
