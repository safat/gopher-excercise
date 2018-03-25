package main

import (
	"github.com/safat/crawler"
	"fmt"
	"strings"
	"regexp"
	"net/http"
	"io/ioutil"
)

const DateRegex string = `\d{4}/\d{2}/\d{2}`

func main() {
	baseUrl := "https://bdnews24.com/"
	visiteMap := make(map[string]bool)

	dateRegex := regexp.MustCompile(DateRegex)

	contentLinks := make([]string, 0)
	categoryLinks := make([]string, 0)

	anchorList := crawler.Crawl(baseUrl)
	visiteMap[baseUrl] = true

	for _, link := range anchorList {
		_, ok := visiteMap[link.Link]

		if ok || !strings.HasPrefix(link.Link, baseUrl) || dateRegex.MatchString(link.Link) {
			continue
		}

		categoryLinks = append(categoryLinks, link.Link)
		contentLinkSlice := findContentLinks(baseUrl, link.Link)

		contentLinks = append(contentLinks, contentLinkSlice...)
	}

	contentMap := make(map[string]string)

	for _, contentLink := range contentLinks {
		response, _ := http.Get(contentLink)
		responseBytes, _ := ioutil.ReadAll(response.Body)
		contentMap[contentLink] = string(responseBytes)

		fmt.Println(contentMap[contentLink])
	}

	fmt.Println(contentMap)
}

func findContentLinks(baseUrl string, categoryLink string) []string {
	dateRegex := regexp.MustCompile(DateRegex)
	contentLinks := make([]string, 0)

	anchorList := crawler.Crawl(categoryLink)

	for _, link := range anchorList {
		if strings.HasPrefix(link.Link, baseUrl) && dateRegex.MatchString(link.Link) {
			contentLinks = append(contentLinks, link.Link)
			continue
		}
	}

	if len(contentLinks) == 0 {
		contentLinks = append(contentLinks, categoryLink)
	}

	fmt.Println(categoryLink, len(contentLinks))

	return contentLinks
}
