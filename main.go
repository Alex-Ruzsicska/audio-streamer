package main

import (
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

func main() {

	database, _ := sql.Open("sqlite3", "db/database.db")
	
	var songName string
	database.QueryRow("SELECT url FROM audio where id = 1").Scan(&songName)

	fmt.Println(songName)

	// configure the songs directory name and port
	const port = 8080

	// add a handler for the song files
	http.Handle("/", addHeaders(http.FileServer(http.Dir("mock/audios"))))
	fmt.Printf("Starting server on %v\n", port)
	log.Printf("Serving %s on HTTP port: %v\n", songName, port)

	// serve and log errors
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

// addHeaders will act as middleware to give us CORS support
func addHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}

//ffmpeg -i BachGavotteShort.mp3 -c:a libmp3lame -b:a 128k -map 0:0 -f segment -segment_time 10 -segment_list outputlist.m3u8 -segment_format mpegts output%03d.ts