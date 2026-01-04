package main

import (
	"embed"
	"encoding/json"
	"harmonycook/music"
	"io/fs"
	"log"
	"net/http"
)

type APIResponse struct {
	ErCode int `json:"erCode"`
	Data   any `json:"data"`
}

var GUI_BUILD_DIR string = "gui-dist"

//go:embed uiweb
var uiweb embed.FS

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	requestedFilepath := GUI_BUILD_DIR + r.URL.Path

	// 	if _, err := os.Stat(requestedFilepath); err == nil {
	// 		http.ServeFile(w, r, requestedFilepath)
	// 		return
	// 	}
	// 	http.ServeFile(w, r, GUI_BUILD_DIR+"/index.html")
	// })

	fsys, _ := fs.Sub(uiweb, "uiweb/dist")
	http.Handle("/", http.FileServer(http.FS(fsys)))

	http.HandleFunc("/api/suggestchords", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

		var reqBody struct {
			Notes []string `json:"notes"`
		}

		if err := json.NewDecoder(req.Body).Decode(&reqBody); err != nil {
			json.NewEncoder(w).Encode(APIResponse{ErCode: 1, Data: nil})
			return
		}

		resData := map[int][]string{}

		notes := music.NewNotes(reqBody.Notes)
		chords := music.SuggestChords(notes)
		for k, v := range chords {
			resData[k] = music.FormatChords(v)
		}

		json.NewEncoder(w).Encode(APIResponse{ErCode: 0, Data: resData})

	})

	log.Println("App is running at http://localhost:5000")
	http.ListenAndServe(":5000", nil)
}
