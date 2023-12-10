package main

import (
	"fmt"
	"log"
	"strings"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
)

func main() {

	converter := md.NewConverter("", true, nil)

	converter.AddRules(
		md.Rule{
			Filter: []string{"del", "s", "strike"},
			Replacement: func(content string, selec *goquery.Selection, opt *md.Options) *string {
				// You need to return a pointer to a string (md.String is just a helper function).
				// If you return nil the next function for that html element
				// will be picked. For example you could only convert an element
				// if it has a certain class name and fallback if not.
				content = strings.TrimSpace(content)
				return md.String("~" + content + "~")
			},
		},
		// more rules
	)

	html := `<strike>strike me</strike>
	<strong>Important</strong>`

	markdown, err := converter.ConvertString(html)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("md ->", markdown)
}
