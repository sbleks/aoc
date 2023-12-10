package scrape

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func ExampleScrape() {
	// Request the HTML page.
	res, err := http.Get("http://metalsucks.net")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find(".left-content article .post-title").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		title := s.Find("a").Text()
		fmt.Printf("Review %d: %s\n", i, title)
	})
}

func ScrapeDays(year string) ([][]string, error) {
	if year == "" {
		year = "2023"
	}

	return scrapeDays(year)
}

func scrapeDays(year string) (out [][]string, err error) {
	// Request the HTML page.
	baseURL := "http://adventofcode.com"
	res, err := http.Get(baseURL + "/" + year)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	doc.Find("pre.calendar a").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the title
		link, exists := s.Attr("href")
		if exists {
			parts := strings.Split(link, "/")
			dayNum := parts[len(parts)-1]
			day := []string{dayNum, baseURL + link}
			out = append(out, day)
			fmt.Printf("a tag %d: Day %s - %s\n", i, dayNum, link)
		}
	})
	return
}

func main() {

}
