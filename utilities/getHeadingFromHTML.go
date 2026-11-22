package utilities

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func GetHeadingFromHTML(html string) (string, error) {

	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal(err)
	}
	heading := doc.Find("h1").Text()
	//fmt.Println(heading)
	return heading, err
}
