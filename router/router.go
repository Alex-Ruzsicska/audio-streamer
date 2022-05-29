package router

import ( 
	"github.com/gorilla/mux"
	"audio-streamer/controller"
	"net/http"
)

//// addHeaders will act as middleware to give us CORS support
func addHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}

func Router() *mux.Router{
	router := mux.NewRouter()

	//Folder
	router.HandleFunc("/folders", controller.GetAllFolders).Methods("GET")
	router.HandleFunc("/folders/genre/{id}", controller.GetFoldersByGenre).Methods("GET")

	//Audio
	router.HandleFunc("/audios", controller.GetAllAudios).Methods("GET")
	router.HandleFunc("/audios/folder/{id}", controller.GetAudiosByFolder).Methods("GET")

	//Genre
	router.HandleFunc("/genres", controller.GetAllGenres).Methods("GET")
	
	//Author
	router.HandleFunc("/authors", controller.GetAllAuthors).Methods("GET")
	
	//Stream
	router.HandleFunc("/stream/audio", controller.StreamAudio).Methods("GET")
	fs := addHeaders(http.FileServer(http.Dir("./mock/")))
	router.PathPrefix("/stream/").Handler(http.StripPrefix("/stream/", fs))


	return router
}