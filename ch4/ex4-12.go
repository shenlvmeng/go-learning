package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Comic struct {
	Num        int
	Transcript string
	Title      string
	Img        string
}

const (
	commicPrefix = "https://xkcd.com/"
	commicSuffix = "/info.0.json"
	maxLimit     = 100 // 避免加载时间太长，仅示意
)

var xkcdDb []*Comic

func main() {
	fmt.Println("Fetching...")
	i := 1
	for i <= maxLimit {
		comic, err := SearchComic(i)
		if err != nil {
			break
		}
		xkcdDb = append(xkcdDb, comic)
		i++
		fmt.Printf("Finish %d\n", i)
	}
	fmt.Println("Fetching end.")
	fmt.Println("Start searching")
	Search()
}

func SearchComic(index int) (*Comic, error) {
	res, err := http.Get(commicPrefix + strconv.Itoa(index) + commicSuffix)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("search error: %s\n", res.Status)
	}

	var jsonRes Comic
	if err := json.NewDecoder(res.Body).Decode(&jsonRes); err != nil {
		res.Body.Close()
		return nil, err
	}
	return &jsonRes, nil
}

func Search() {
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	in.Scan()
	indexStr := in.Text()

	id, err := strconv.Atoi(indexStr)
	if err != nil {
		fmt.Println("Start new search.")
		Search()
		return
	}

	if id >= len(xkcdDb) {
		fmt.Println("Out fo db range")
		fmt.Println("Start new search.")
		Search()
		return
	}

	comic := xkcdDb[id]
	fmt.Printf("title\t\turl\t\ttranscript\n")
	fmt.Printf("%s\t\t%s\t\t%s", comic.Title, comic.Img, comic.Transcript)
	fmt.Println("Start new search.")
	Search()
}
