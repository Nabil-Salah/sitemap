package main

import (
	"flag"
	"net/http"
)

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "the url that you want to build a sitemap for")
	//maxDepth := flag.Int("depth", 10, "the maximum number of links deep to traverse")
	flag.Parse()
	resp, err := http.Get(*urlFlag)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
}
