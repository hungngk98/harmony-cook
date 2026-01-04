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

type APIResponse struct {
	ErCode  int    `json:"erCode"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

//go:embed uiweb
var uiwebEmbedFS embed.FS

func main() {
	flag.Parse()

	if *envFlag == "prod" {
		uiwebDist, _ := fs.Sub(uiwebEmbedFS, "uiweb/dist")

		muxUIWeb := http.NewServeMux()

		muxUIWeb.Handle("/", http.FileServer(http.FS(uiwebDist)))

		go func() {
			fmt.Println("Use app at http://localhost:3000")
			log.Fatal(http.ListenAndServe(":3000", muxUIWeb))
		}()
	}

	muxServer := http.NewServeMux()

	muxServer.HandleFunc("/api/suggestchords", func(w http.ResponseWriter, req *http.Request) {
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

		notes, err := music.NewNotes(reqBody.Notes)
		if err != nil {
			json.NewEncoder(w).Encode(APIResponse{ErCode: 1, Message: "invalid notes"})
			return
		}

		chords := music.SuggestChords(notes)
		for k, v := range chords {
			resData[k] = music.FormatChords(v)
		}

		json.NewEncoder(w).Encode(APIResponse{ErCode: 0, Data: resData})

	})

	go func() {
		log.Println("Server is running")
		log.Fatal(http.ListenAndServe(":5000", muxServer))
	}()

	select {}
}
