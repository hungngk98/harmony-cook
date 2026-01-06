package server

import (
	"fmt"
	"harmonycook/embedstatic"
	"log"
	"net/http"
)

func RunAPIServer() {

	muxServer := http.NewServeMux()

	muxServer.HandleFunc("/api/suggestchords", SuggestChordsHandler)
	muxServer.HandleFunc("/api/suggesttones", SuggestTonesHandler)

	log.Fatal(http.ListenAndServe(":5000", muxServer))
}

func RunFileServer() {
	uiwebFS, _ := embedstatic.GetFS("uiweb")

	muxUIWeb := http.NewServeMux()

	muxUIWeb.Handle("/", http.FileServer(http.FS(uiwebFS)))

	fmt.Println("Use app at http://localhost:3000")

	log.Fatal(http.ListenAndServe(":3000", muxUIWeb))
}
