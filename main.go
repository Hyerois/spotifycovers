package main

import (
	"fmt"
	"strconv"
	"strings"
)

func exit(err error) { //DISPLAYING ERROR THEN QUITTING PROGRAM
	fmt.Println("Erreur durant l'exécution du programme:")
	fmt.Println(err)
	fmt.Println("Appuyer sur n'importe quelle touche pour finir l'exécution du programme.")
	fmt.Scanln()
	return
}

func main() {
	token, err := retrieveToken() //RETRIEVE TOKEN FOR API USE
	if err != nil {
		exit(err)
	}

	fmt.Println(`
	Pour trouver l'id d'un album ou d'un artiste:
	Il suffit de se rendre sur la page Spotify de l'artiste en question puis de copier la chaîne de caractère qui se trouve dans l'URL.
	Exemple: "https://open.spotify.com/artist/1Xyo4u8uXC1ZmMpatF05PJ" est l'adresse de la page de The Weeknd.
	Ici l'id est: "1Xyo4u8uXC1ZmMpatF05PJ" (sans les guillemets).
	Entrer "1" pour le lien de la photo d'un artiste
	Entrer "2" pour le lien de la cover d'un album
	Entrer "3" pour finir l'exécution du programme
	
	Voulez-vous enregistrer les résultats dans un fichier local à la racine du programme ? O/n
	`)
	var saveChoice string
	var doSave bool
	fmt.Scanln(&saveChoice)
	doSave = strings.ToLower(saveChoice) == "o" || strings.ToLower(saveChoice) == ""

	for true {
		var choice string
		fmt.Print(`Choix: `)
		fmt.Scanln(&choice) //USER INPUT FOR CHOICE : {"1","2","3"}

		if choice == "1" { //IF CHOICE IS 1 RETRIEVE ARTIST PROFILE PICTURE WITH HIS SPOTIFY ID
			var artistId string
			fmt.Println("Entrer l'id Spotify de l'artiste.")
			fmt.Print("Id: ")
			fmt.Scanln(&artistId)                          //USER INPUT FOR ID
			artist, err := retrieveArtist(artistId, token) //RETRIEVE ARTIST INFORMATIONS

			if err != nil {
				if err.Error() == "404" {
					fmt.Println("L'id n'existe pas.")
					continue
				} else {
					exit(err)
				}
			}

			if doSave {
				writeArtistToFile(artist)
			}

			fmt.Println("Images pour l'artiste " + artist.Name + " :")
			for _, image := range artist.Images { //PRINTING SIZE AND LINK FOR EVERY PICTURES
				fmt.Println(strconv.Itoa(image.Width) + "x" + strconv.Itoa(image.Height) + " : " + image.Url)
			}
			fmt.Println()

		} else if choice == "2" { //SAME AS BEFORE
			var albumId string
			fmt.Println("Entrer l'id Spotify de l'album.")
			fmt.Print("Id: ")
			fmt.Scanln(&albumId)
			album, err := retrieveAlbum(albumId, token)

			if err != nil {
				if err.Error() == "404" {
					fmt.Println("L'id n'existe pas.")
					continue
				} else {
					exit(err)
				}
			}

			if doSave {
				writeAlbumToFile(album)
			}

			fmt.Println("Images pour l'album " + album.Name + " :")
			for _, image := range album.Images {
				fmt.Println(strconv.Itoa(image.Width) + "x" + strconv.Itoa(image.Height) + " : " + image.Url)
			}
			fmt.Println()

		} else if choice == "3" {
			return

		} else {
			fmt.Println("Mauvais choix.")
		}
	}
}
