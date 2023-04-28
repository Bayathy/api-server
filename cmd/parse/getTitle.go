package parse

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

func GetTitle(url string) (string, error) {

	res, err := http.Get(url)

	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	title := doc.Find("title").Text()

	return title, nil
}
