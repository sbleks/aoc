package parsemd

import (
	"aoc2md/scrape"
	"fmt"
	"log"
	"os"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
)

func New() *md.Converter {
	converter := md.NewConverter("", true, nil)

	converter.AddRules(
		md.Rule{
			Filter: []string{"code"},
			Replacement: func(content string, selec *goquery.Selection, opt *md.Options) *string {
				// You need to return a pointer to a string (md.String is just a helper function).
				// If you return nil the next function for that html element
				// will be picked. For example you could only convert an element
				// if it has a certain class name and fallback if not.
				// log.Printf("children text: %s, text: %s\n", selec.Children().Text(), selec.Text())
				if selec.Children().Text() == selec.Text() {
					content = strings.TrimSpace(content)
					return md.String("_`" + selec.Text() + "`_")
				} else {
					return md.String("`" + content + "`")
				}
			},
		},
		md.Rule{
			Filter: []string{"span"},
			Replacement: func(content string, selec *goquery.Selection, opt *md.Options) *string {
				html, err := goquery.OuterHtml(selec)
				if err != nil {
					return md.String(content)
				}

				return md.String(html)
			},
		},
		// more rules
	)
	return converter
}

func ConvertDay(selection *goquery.Selection) (string, error) {
	converter := New()
	return converter.Convert(selection), nil

}

func RunConvert() {
	baseURL := "https://adventofcode.com"
	year := "2023"
	day := "4"
	dayURL := "/" + year + "/day/" + day
	selection, title, err := scrape.ScrapeDay(dayURL)
	if err != nil {
		log.Panicf("Could not scrape day: %v\n", err)
	}

	markdown, err := ConvertDay(selection)
	if err != nil {
		log.Panicf("Could not convert day: %v\n", err)
	}

	header := fmt.Sprintf("# %s\n\n[%s](%s)\n\n## Description\n\n### Part One\n\n", title, baseURL+dayURL, baseURL+dayURL)
	markdown = header + markdown
	log.Printf("Markdown:\n%s\n", markdown)

	err = os.MkdirAll("../"+year+"/day"+day+"/", os.FileMode(int(0777)))
	if err != nil {
		log.Panicf("Could not make dir: %v\n", err)
	}

	os.Chdir("../" + year + "/day" + day + "/")
	os.WriteFile("README.md", []byte(markdown), os.FileMode(int(0777)))

	starterText := `package main

import (
	"aocInput"
	"log"
	"strconv"
	"strings"
)

func part1(lines []string) (sum int) {
	return 0
}


func part1(lines []string) (sum int) {
	return 0
}

func main() {

	lines, err := input.GetInputLines("./example.txt")
	if err != nil {
		log.Fatalf("Could not read file: %v", err)
	}

	// sum := part1(lines)
	// sum := part2(lines)

	log.Printf("Sum is: %v", sum)
}
`

	os.WriteFile("main.go", []byte(starterText), os.FileMode(int(0777)))
}
