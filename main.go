package main

import (
	"flag"
	"harmonycook/server"
)

var devFlag = flag.Bool("dev", false, "Running environment")

func main() {
	flag.Parse()

	if *devFlag == false {
		go func() { server.RunFileServer() }()
	}

	server.RunAPIServer()
}
