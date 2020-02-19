package main

type Token struct {
	String    string `json:"access_token"`
	TokenType string `json:"token_type"`
	ExpiresIn int    `json:"expires_in"`
	Scope     string `json:"scope"`
}

type SimplifiedAlbum struct {
	AlbumType string  `json:"album_type"`
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	Images    []Image `json:"images"`
}

type ArtistAlbumsResponse struct {
	Items []SimplifiedAlbum `json:"items"`
}

type Artist struct {
	Name   string  `json:"name"`
	Id     string  `json:"id"`
	Images []Image `json:"images"`
}

type AlbumTracksResponse struct {
	Items []AlbumTracks `json:"items"`
}

type AlbumTracks struct {
	Artists []Artist `json:"artists"`
	Id      string   `json:"id"`
	Name    string   `json:"name"`
}

type Image struct {
	Height int    `json:"height"`
	Width  int    `json:"width"`
	Url    string `json:"url"`
}
