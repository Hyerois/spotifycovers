package main

import (
	"encoding/json"
	"os"
)

func writeArtistToFile(artist Artist) {
	f, err := os.Create(artist.Name + ".json")
	str, err := json.MarshalIndent(artist, "", " ")
	_, err = f.WriteString(string(str))
	f.Close()
	if err != nil {
		exit(err)
	}
}

func writeAlbumToFile(album SimplifiedAlbum) {
	f, err := os.Create(album.Name + ".json")
	str, err := json.MarshalIndent(album, "", " ")
	_, err = f.WriteString(string(str))
	f.Close()
	if err != nil {
		exit(err)
	}
}
