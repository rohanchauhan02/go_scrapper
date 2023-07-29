package main

import (
	"encoding/csv"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func gethtml(url string) *http.Response {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	return resp
}

func main() {
	url := "https://www.ebay.com/sch/i.html?_from=R40&_trksid=p2380057.m570.l1313&_nkw=iphone&_sacat=0"

	var previousURL string

	for i := 0; i < 10; i++ {
		resp := gethtml(url)
		defer resp.Body.Close()

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			panic(err)
		}
		scrapDocument(doc)

		href, _ := doc.Find("nav.pagination>a.pagination__next").Attr("href")
		if url == previousURL {
			break
		}
		previousURL = url
		url = href
	}

}

func scrapDocument(doc *goquery.Document) {
	doc.Find("ul.srp-results>li.s-item").Each(func(i int, s *goquery.Selection) {
		a := s.Find("a.s-item__link")
		// Title
		title := strings.TrimSpace(a.Text())
		// Link
		link, _ := a.Attr("href")
		// Price
		price := s.Find("span.s-item__price").Text()
		scrapData := []string{title, link, price}
		// fmt.Println(scrapData)
		writeToCSV(scrapData)
		writeToXML(scrapData)
		writeToJSON(scrapData)
	})
}

func writeToCSV(data []string) {
	fileName := "data.csv"

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	err = writer.Write(data)
	if err != nil {
		panic(err)
	}
}

func writeToXML(data []string) {
	fileName := "data.xml"

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	err = writer.Write(data)
	if err != nil {
		panic(err)
	}
}

func writeToJSON(data []string) {
	fileName := "data.json"

	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	err = writer.Write(data)
	if err != nil {
		panic(err)
	}
}

