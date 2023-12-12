package scrape

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func stripHyphens(str string) string {
	r := strings.NewReplacer("--- ", "", " ---", "")
	return r.Replace(str)
}

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
			// fmt.Printf("a tag %d: Day %s - %s\n", i, dayNum, link)
		}
	})
	return
}

func ScrapeDay(url string) (*goquery.Selection, string, error) {
	baseURL := "http://adventofcode.com"
	cookie := http.Cookie{Name: "session", Value: "53616c7465645f5f73a7044c36ed47476f83951350c98741621a00141d008a7e64f8b8da15349b02a9f455e20309bc08d708a56b8edba9d18b975584d35e28f5"}
	req, err := http.NewRequest(http.MethodGet, baseURL+url, http.NoBody)
	if err != nil {
		log.Fatal(err)
	}
	req.AddCookie(&cookie)

	res, err := http.DefaultClient.Do(req)
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

	title := ""
	// Return the selection to be used in the parser
	html := doc.Find("article").Each(func(i int, s *goquery.Selection) {
		h2 := s.Find("h2")
		if i == 0 {
			title = stripHyphens(h2.Text())
			h2.Remove()
			return
		}
		h2.ReplaceWithHtml("<h3>" + stripHyphens(h2.Text()) + "</h3>")

	})

	// text, _ := html.Html()

	// log.Printf("%v\n", text)

	return html, title, nil
}
