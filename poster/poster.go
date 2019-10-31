package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Movie struct with main information about
type Movie struct {
	Title      string
	Year       string
	Runtime    string
	Genre      string
	Director   string
	Actors     string
	ImdbRating string
}

const url = "http://www.omdbapi.com/?apikey=422b5472&t="

func main() {
	title := os.Args[1]
	var movie Movie

	if title == "" {
		fmt.Fprintf(os.Stderr, "Err: poster [title]\n")
	}

	resp, err := http.Get(url + title)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s\n", err)
	}

	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s\n", err)
	}

	if movie != (Movie{}) {
		fmt.Printf("%s\t%s\t%s\n%s\n%s\n%s\n%s rate\n",
			movie.Title, movie.Year, movie.Runtime,
			movie.Genre,
			movie.Director,
			movie.Actors,
			movie.ImdbRating,
		)
	} else {
		fmt.Printf("No information\n")
	}
}
