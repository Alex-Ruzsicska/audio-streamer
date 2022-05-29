package controller

import (
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"database/sql"
	"log"
	"audio-streamer/model"
	"encoding/json"
)

func GetAllAudios(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all audios")
	var audios []model.Audio
	database, _ := sql.Open("sqlite3", "db/database.db")
	rows, err := database.Query("SELECT * FROM audio")

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var audio model.Audio
		err := rows.Scan(&audio.Id, &audio.FolderId, &audio.Name, &audio.Url)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(audio.Id, audio.FolderId, audio.Name, audio.Url)
		audios = append(audios, audio)
	}

	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(audios)
}

func GetAllFolders(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all folders")
	var folders []model.Folder
	database, _ := sql.Open("sqlite3", "db/database.db")
	rows, err := database.Query("SELECT * FROM folder")

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var folder model.Folder
		err := rows.Scan(&folder.Id, &folder.Name, &folder.CoverUrl)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(folder.Id, folder.Name, folder.CoverUrl)
		folders = append(folders, folder)
	}

	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(folders)
}

func GetAudio(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r);
	audioId := params["id"]
	fmt.Println("Get audio", audioId)

	var audio model.Audio
	database, _ := sql.Open("sqlite3", "db/database.db")
	row := database.QueryRow("SELECT * FROM audio where id = ?", audioId)

	row.Scan(&audio.Id, &audio.FolderId, &audio.Name, &audio.Url)
	log.Println(audio.Id, audio.FolderId, audio.Name, audio.Url)

	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(audio)
}

func GetAllGenres(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all genres")
	var genres []model.Genre
	database, _ := sql.Open("sqlite3", "db/database.db")
	rows, err := database.Query("SELECT * FROM genre")

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var genre model.Genre
		err := rows.Scan(&genre.Id, &genre.Name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(genre.Id, genre.Name)
		genres = append(genres, genre)
	}

	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(genres)
}

func GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all authors")
	var authors []model.Author
	database, _ := sql.Open("sqlite3", "db/database.db")
	rows, err := database.Query("SELECT * FROM author")

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var author model.Author
		err := rows.Scan(&author.Id, &author.Name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(author.Id, author.Name)
		authors = append(authors, author)
	}

	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(authors)
}

func GetFoldersByGenre(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r);
	genreId := params["id"]

	fmt.Println("Get folders by genre", genreId)

	var folders []model.Folder
	database, _ := sql.Open("sqlite3", "db/database.db")
	rows, err := database.Query("SELECT id, name, cover_url FROM folder left JOIN genre_folder ON genre_folder.folder_id = folder.id WHERE genre_folder.genre_id = ?", genreId)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var folder model.Folder
		err := rows.Scan(&folder.Id, &folder.Name, &folder.CoverUrl)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(folder.Id, folder.Name, folder.CoverUrl)
		folders = append(folders, folder)
	}

	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(folders)
}

func GetAudiosByFolder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r);
	folderId := params["id"]

	fmt.Println("Get audios by folder", folderId)

	var audios []model.Audio
	database, _ := sql.Open("sqlite3", "db/database.db")
	rows, err := database.Query("SELECT * FROM audio WHERE folder_id = ?", folderId)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var audio model.Audio
		err := rows.Scan(&audio.Id, &audio.FolderId, &audio.Name, &audio.Url)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(audio.Id, audio.FolderId, audio.Name, audio.Url)
		audios = append(audios, audio)
	}

	w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(audios)
}

func StreamAudio(w http.ResponseWriter, r *http.Request){
	addHeaders(http.FileServer(http.Dir("mock/audios")))
	h := http.FileServer(http.Dir("mock/audios"))
	w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
}

// addHeaders will act as middleware to give us CORS support
func addHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}

