package server

import (
	"encoding/json"
	"harmonycook/music"
	"net/http"
)

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
