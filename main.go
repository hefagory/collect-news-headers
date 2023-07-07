package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://example.com/news" // Replace with the actual URL of the news page

	// Make an HTTP GET request to the news page
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Parse the HTML document
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find and collect the headers of news articles
	headers := make([]string, 0)
	document.Find(".news-header").Each(func(index int, element *goquery.Selection) {
		header := element.Text()
		headers = append(headers, header)
	})

	// Print the collected headers
	for _, header := range headers {
		fmt.Println(header)
	}
}
