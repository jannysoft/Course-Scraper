package util

import (
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

func ScrapeHashTags() (string) {

	// Open the submitted to standard out url
	response, _ := openRequest("http://tweeplers.com/hashtags/?cc=WORLD")

	var hashTags string
	defer response.Body.Close()
	// Open the returned response.Body with goquery and assign it to var doc
	doc, _ := goquery.NewDocumentFromReader(response.Body)

	hashtag1 := doc.Find("#item_u_1").Text()
	hashtag2 := doc.Find("#item_u_2").Text()
	hashtag3 := doc.Find("#item_u_3").Text()

	hashTags = hashtag1 + " " + hashtag2 + " " + hashtag3

	return hashTags
}

func openRequest(url string) (*http.Response, error) {
	browser := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")
	response, err := browser.Do(request)

	if err != nil {
		return nil, err
	}
	return response, nil
}
