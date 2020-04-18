package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type Movie struct {
	Title      string
	Year       string
	Rated      string
	Director   string
	Poster     string
	IMDBRating string `json:"imdbRating"`
	IMDBID     string `json: imdbID`
}

const OMDBURL = "http://www.omdbapi.com/?apikey=4dd25ec6"

func main() {
	fmt.Println("Print movie name to search")
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	name := in.Text()

	res, err := http.Get(OMDBURL + "&t=" + name)
	if err != nil {
		fmt.Printf("fetch error: %v\n", err)
		os.Exit(1)
	}
	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		fmt.Printf("fetch error: %s\n", res.Status)
		os.Exit(1)
	}

	var movie Movie
	if err := json.NewDecoder(res.Body).Decode(&movie); err != nil {
		fmt.Printf("json decode error: %v\n", err)
		os.Exit(1)
	}
	fetchPoster(movie.Poster)
}

func fetchPoster(posterUrl string) error {
	if posterUrl == "" || posterUrl == "N/A" {
		fmt.Println("empty poster url")
		return nil
	}
	res, err := http.Get(posterUrl)
	if err != nil {
		fmt.Printf("fetch error: %v\n", err)
		return err
	}
	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return fmt.Errorf("fetch error: %s\n", res.Status)
	}
	out, err := os.Create("./poster.png")
	if err != nil {
		fmt.Printf("create image error: %v\n", err)
		return err
	}
	if _, err = io.Copy(out, res.Body); err != nil {
		res.Body.Close()
		fmt.Printf("download image error: %v\n", err)
		return err
	}
	return nil
}
