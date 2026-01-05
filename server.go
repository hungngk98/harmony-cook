package main

import (
	"embed"
	"encoding/json"
	"flag"
	"fmt"
	"harmonycook/music"
	"io/fs"
	"log"
	"net/http"
)

var envFlag = flag.String("env", "prod", "Running environment")

//go:embed uiweb
var uiwebFS embed.FS

func main() {
	flag.Parse()

	if *envFlag == "prod" {
		uiwebDistFS, _ := fs.Sub(uiwebFS, "uiweb/dist")

		muxUIWeb := http.NewServeMux()

		muxUIWeb.Handle("/", http.FileServer(http.FS(uiwebDistFS)))

		go func() {
			log.Fatal(http.ListenAndServe(":3000", muxUIWeb))
		}()

		fmt.Println("Use app at http://localhost:3000")
	}

	muxServer := http.NewServeMux()

	muxServer.HandleFunc("/api/suggestchords", SuggestChordsHandler)
	muxServer.HandleFunc("/api/suggesttones", SuggestTonesHandler)

	log.Fatal(http.ListenAndServe(":5000", muxServer))
}

type APIResponse struct {
	ErCode  int    `json:"erCode"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func SuggestChordsHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	var reqBody struct {
		Notes []string `json:"notes"`
	}

	if err := json.NewDecoder(req.Body).Decode(&reqBody); err != nil {
		json.NewEncoder(w).Encode(APIResponse{ErCode: 1, Message: "server error"})
		return
	}

	resData := map[int][]string{}

	chords := music.SuggestChords(music.NewNotes(reqBody.Notes))
	for k, v := range chords {
		resData[k] = music.FormatChords(v)
	}

	json.NewEncoder(w).Encode(APIResponse{ErCode: 0, Data: resData})
}

func SuggestTonesHandler(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	var reqBody struct {
		Notes []string `json:"notes"`
	}

	if err := json.NewDecoder(req.Body).Decode(&reqBody); err != nil {
		json.NewEncoder(w).Encode(APIResponse{ErCode: 1, Message: "server error"})
		return
	}

	resData := map[int][]string{}

	scales := music.SuggestTones(music.NewNotes(reqBody.Notes))
	for k, v := range scales {
		resData[k] = music.FormatScales(v)
	}

	json.NewEncoder(w).Encode(APIResponse{ErCode: 0, Data: resData})
}
