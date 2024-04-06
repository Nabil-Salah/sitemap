package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"net/url"
	"os"
	link "sitemap"
	"strings"
)

const xmlns = "http://www.sitemaps.org/schemas/sitemap/0.9"

type loc struct {
	Value string `xml:"loc"`
}

type urlset struct {
	Urls  []loc  `xml:"url"`
	Xmlns string `xml:"xmlns,attr"`
}

func main() {
	urlFlag := flag.String("url", "https://www.sitemaps.org/schemas/sitemap/", "the url that you want to build a sitemap for")
	maxDepth := flag.Int("depth", 10, "the maximum number of links deep to traverse")
	flag.Parse()
	pages := bfs(urlFlag, *maxDepth)
	toXml := urlset{
		Xmlns: xmlns,
	}
	for _, page := range pages {
		toXml.Urls = append(toXml.Urls, loc{page})
	}
	fmt.Print(xml.Header)
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("", "  ")
	if err := enc.Encode(toXml); err != nil {
		panic(err)
	}
	fmt.Println()
}
func bfs(urlStr *string, maxDepth int) []string {
	seen := make(map[string]int)
	var q []string
	idx := 1
	nq := []string{
		*urlStr,
	}
	for i := 0; i <= maxDepth; i++ {
		q, nq = nq, []string{}
		if len(q) == 0 {
			break
		}
		for _, url := range q {
			if val, ok := seen[url]; val != 0 && ok {
				continue
			}
			seen[url] = idx
			idx += 1
			for _, link := range get(&url) {
				nq = append(nq, link)
			}
		}
	}
	ret := make([]string, len(seen))
	for url, idx := range seen {
		ret[idx-1] = url
	}
	return ret
}
func get(urlStr *string) []string {
	resp, err := http.Get(*urlStr)
	if err != nil {
		return []string{}
	}
	defer resp.Body.Close()
	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()
	return filter(hrefs(html.NewTokenizer(resp.Body), base), withPrefix(&base))
}
func withPrefix(pfx *string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, *pfx)
	}
}
func hrefs(r *html.Tokenizer, base string) []string {
	links, _ := link.Parse(r)
	var ret []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			ret = append(ret, base+l.Href)
		case strings.HasPrefix(l.Href, "http"):
			ret = append(ret, l.Href)
		}
	}

	return ret
}

func filter(links []string, keepFn func(string) bool) []string {
	var ret []string
	for _, link := range links {
		if keepFn(link) {
			ret = append(ret, link)
		}
	}
	return ret
}
