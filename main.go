package main

import (
	"fmt"
	"log"
	"strconv"
)

const baseApi = "https://accounts.spotify.com/api/token"
const apiKey = "NzVhNjJkODgyMDZlNDQ1Y2EyNDQzMDBiNDk5ZWFjMzA6MzkzMDkwMzNkZjQ1NDAyMGI3NGRhZDAxYjI4MDJjYWI="

func handleError(err error) {
	if err != nil {
		fmt.Println("Erreur durant l'exécution du programme:")
		fmt.Println(err)
		fmt.Println("Appuyer sur n'importe quelle touche pour finir l'exécution du programme.")
		fmt.Scanln()
		log.Fatal()
	}
}

func main() {
	token, err := retrieveToken()
	handleError(err)
	var choice string
	fmt.Println(`Entrer "1" pour le lien de la photo d'un artiste`)
	fmt.Println(`Entrer "2" pour le lien de la cover d'un album`)
	for choice != "1" && choice != "2" {
		fmt.Print(`Choix: `)
		fmt.Scanln(&choice)
	}
	if choice == "1" {
		var artistId string
		fmt.Println("Entrer l'id Spotify de l'artiste.")
		fmt.Println(`Pour trouver l'id de l'artiste il suffit de se rendre sur la page Spotify de l'artiste en question puis de copier la chaîne de caractère qui se trouve dans l'URL.`)
		fmt.Println(`Exemple: "https://open.spotify.com/artist/1Xyo4u8uXC1ZmMpatF05PJ" est l'adresse de la page de The Weeknd.`)
		fmt.Println(`Ici l'id est: "1Xyo4u8uXC1ZmMpatF05PJ" (sans les guillemets).`)
		fmt.Print("Id: ")
		fmt.Scanln(&artistId)
		artist, err := retrieveArtist(artistId, token)
		handleError(err)
		fmt.Println("Images pour l'artiste " + artist.Name + " :")
		for _, image := range artist.Images {
			fmt.Println(strconv.Itoa(image.Width) + "x" + strconv.Itoa(image.Height) + " : " + image.Url)
		}
		fmt.Println("Appuyez sur n'importe quelle touche pour continuer.")
		fmt.Scanln()
	}
	if choice == "2" {
		var albumId string
		fmt.Println("Entrer l'id Spotify de l'album.")
		fmt.Println(`Pour trouver l'id de l'album il suffit de se rendre sur la page Spotify de l'album en question puis de copier la chaîne de caractère qui se trouve dans l'URL.`)
		fmt.Println(`Exemple: "https://open.spotify.com/album/4qZBW3f2Q8y0k1A84d4iAO" est l'adresse de la page de l'album "My Dear Melancholy".`)
		fmt.Println(`Ici l'id est: "4qZBW3f2Q8y0k1A84d4iAO" (sans les guillemets).`)
		fmt.Print("Id: ")
		fmt.Scanln(&albumId)
		album, err := retrieveAlbum(albumId, token)
		handleError(err)
		fmt.Println("Images pour l'album " + album.Name + " :")
		for _, image := range album.Images {
			fmt.Println(strconv.Itoa(image.Width) + "x" + strconv.Itoa(image.Height) + " : " + image.Url)
		}
		fmt.Println("Appuyez sur n'importe quelle touche pour continuer.")
		fmt.Scanln()
	}
	for true {
		var choice string
		fmt.Println(`Entrer "1" pour le lien de la photo d'un artiste`)
		fmt.Println(`Entrer "2" pour le lien de la cover d'un album`)
		fmt.Println(`Entrer "3" pour finir l'exécution du programme`)
		for choice != "1" && choice != "2" && choice != "3" {
			fmt.Print(`Choix: `)
			fmt.Scanln(&choice)
		}
		if choice == "1" {
			var artistId string
			fmt.Println("Entrer l'id Spotify de l'artiste.")
			fmt.Print("Id: ")
			fmt.Scanln(&artistId)
			artist, err := retrieveArtist(artistId, token)
			handleError(err)
			fmt.Println("Images pour l'artiste " + artist.Name + " :")
			for _, image := range artist.Images {
				fmt.Println(strconv.Itoa(image.Width) + "x" + strconv.Itoa(image.Height) + " : " + image.Url)
			}
		}
		if choice == "2" {
			var albumId string
			fmt.Println("Entrer l'id Spotify de l'album.")
			fmt.Print("Id: ")
			fmt.Scanln(&albumId)
			album, err := retrieveAlbum(albumId, token)
			handleError(err)
			fmt.Println("Images pour l'album " + album.Name + " :")
			for _, image := range album.Images {
				fmt.Println(strconv.Itoa(image.Width) + "x" + strconv.Itoa(image.Height) + " : " + image.Url)
			}
		}
		if choice == "3" {
			return
		}
	}
}
