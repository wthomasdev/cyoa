package main

import (
	"flag"
	"fmt"
	"gophercises/cyoa"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the CYOA web application on")
	fileName := flag.String("file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()
	fmt.Printf("Using the story in %s. \n", *fileName)

	f, err := os.Open(*fileName)
	if err != nil {
		log.Fatalf("Could not read the story file: %v", err)
		return
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		log.Fatalf("Could not create story; %v", err)
	}

	h := cyoa.NewHandler(story)
	fmt.Printf("Starting the server on port: %d\n", *port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))

}
