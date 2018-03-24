package main

import (
	"net/http"
	"golang.org/x/net/html"
	"regexp"
	"fmt"
)

type Anchor struct {
	Link string
	Text string
}

func main() {
	url := "https://bangla.bdnews24.com/"
	anchorList := crawl(url)

	fmt.Println("Crawled link size: ", len(anchorList))

}

func crawl(url string) []Anchor {
	urlPrefixRegex, _ := regexp.Compile("^(http|https|ftp).*")
	resp, _ := http.Get(url)

	htmlTagTokenizer := html.NewTokenizer(resp.Body)

	anchorList := make([]Anchor, 0)
Loop:
	for {
		htmlTag := htmlTagTokenizer.Next()

		switch {
		case htmlTag == html.ErrorToken:
			break Loop
		case htmlTag == html.StartTagToken:
			htmlToken := htmlTagTokenizer.Token()

			if htmlToken.Data == "a" {
				anchor, ok := parseAnchorTag(htmlToken, urlPrefixRegex)

				if ok {
					anchorList = append(anchorList, anchor)
				}
			}
		}
	}

	return anchorList
}

func parseAnchorTag(htmlToken html.Token, urlPrefixRegex *regexp.Regexp) (Anchor, bool) {
	for _, attribute := range htmlToken.Attr {
		if attribute.Key == "href" && urlPrefixRegex.MatchString(attribute.Val) {
			return Anchor{attribute.Val, ""}, true
		}
	}

	return Anchor{}, false
}
