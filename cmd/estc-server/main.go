package main

import (
	"flag"
	"log"

	"github.com/achiku/gotodoit/estc"
)

func main() {
	port := flag.String("p", "", "port number")
	flag.Parse()

	if *port != "" {
		s := estc.TestNewServer(estc.DefaultHandlerMap, *port)
		if err := s.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal("service port number is not specified")
	}
}
