package main

import (
	"flag"
	"log"

	"github.com/achiku/gotodoit"
)

var (
	commitHash string
	buildDate  string
	goVersion  string
)

func main() {
	log.Printf("hash: %s\n", commitHash)
	log.Printf("build at %s with %s\n", buildDate, goVersion)

	configFile := flag.String("c", "", "config file path")
	flag.Parse()

	if *configFile != "" {
		gotodoit.RunServer(*configFile)
	} else {
		log.Fatal("config file is not specified")
	}
}
