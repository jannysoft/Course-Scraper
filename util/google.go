package util

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
)

type GoogleResult struct {
	ResultDesc string
}

func GetLinksFromGoogle() (linkList []string) {
	linkList = searchGoogle("https://www.google.com/search?q=%22/%3FcouponCode%3D%22&num=100&source=lnt&tbs=qdr:h2")
	return
}

func searchGoogle(googleSearch string) (scrapedList []string) {
	scraped, _ := googleScrape(googleSearch)

	for _, results := range scraped {
		getStrings := strings.Split(results.ResultDesc, " ")

		for _, link := range getStrings {
			if strings.Contains(link, "/?couponCode=") {
				if strings.HasPrefix(link, "https://www.udemy.") {
					if strings.HasSuffix(link, ".") {
						link = strings.TrimSuffix(link, ".")
						scrapedList = append(scrapedList, link)
					} else if strings.HasSuffix(link, " ...") {
						link = strings.TrimSuffix(link, " ...")
						scrapedList = append(scrapedList, link)
					} else if strings.HasSuffix(link, " ..") {
						link = strings.TrimSuffix(link, " ..")
						scrapedList = append(scrapedList, link)
					} else if strings.HasSuffix(link, ";") {
						scrapedList = append(scrapedList, link)
					} else {
						scrapedList = append(scrapedList, link)
					}
				} else if strings.HasPrefix(link, "com") {
					link = "https://www.udemy." + link
					if strings.HasSuffix(link, ".") {
						link = strings.TrimSuffix(link, ".")
						scrapedList = append(scrapedList, link)
					} else if strings.HasSuffix(link, " ...") {
						link = strings.TrimSuffix(link, " ...")
						scrapedList = append(scrapedList, link)
					} else if strings.HasSuffix(link, " ..") {
						link = strings.TrimSuffix(link, " ..")
						scrapedList = append(scrapedList, link)
					} else if strings.HasSuffix(link, ";") {
						scrapedList = append(scrapedList, link)
					} else {
						scrapedList = append(scrapedList, link)
					}
				} else {
					fmt.Println("no link found")
				}
			}
		}
	}
	return
}

func googleScrape(googleSearch string) ([]GoogleResult, error) {
	res, err := googleRequest(googleSearch)
	if err != nil {
		return nil, err
	}
	scrapes, err := googleResultParser(res)
	if err != nil {
		return nil, err
	} else {
		return scrapes, nil
	}
}

func googleRequest(searchURL string) (*http.Response, error) {

	baseClient := &http.Client{}

	req, _ := http.NewRequest("GET", searchURL, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")

	res, err := baseClient.Do(req)

	if err != nil {
		return nil, err
	} else {
		return res, nil
	}
}

func googleResultParser(response *http.Response) ([]GoogleResult, error) {

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}

	var results []GoogleResult
	sel := doc.Find("div.g")
	for i := range sel.Nodes {
		item := sel.Eq(i)
		descTag := item.Find("span.st")
		desc := descTag.Text()
		result := GoogleResult{
			desc,
		}
		results = append(results, result)
	}
	return results, err
}
