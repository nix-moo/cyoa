package main

import (
	"cyoa"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	port := flag.Int("port", 3000, "the port to start web app on")
	file := flag.String("file", "gopher.json", "the JSON file containing the story text")
	flag.Parse()
	fmt.Printf("Using the story in %s.\n", *file)

	f, err := os.Open(*file)
	if err != nil {
		fmt.Printf("Something went wrong")
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}

	tpl := template.Must(template.New("").Parse("Hello workld!"))
	h := cyoa.NewHandler(story, cyoa.WithTemplate(tpl))
	fmt.Printf("Starting server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
