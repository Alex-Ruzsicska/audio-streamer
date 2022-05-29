package main

import (
	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
	//"database/sql"
	"fmt"
	"log"
	"net/http"
	"audio-streamer/router"

)

func main() {
	r := router.Router();

	fmt.Println("Starting server configuration")
	http.ListenAndServe(":4000", r)
	log.Fatal(http.ListenAndServe(":4000", r))
}

//ffmpeg -i BachGavotteShort.mp3 -c:a libmp3lame -b:a 128k -map 0:0 -f segment -segment_time 10 -segment_list outputlist.m3u8 -segment_format mpegts output%03d.ts