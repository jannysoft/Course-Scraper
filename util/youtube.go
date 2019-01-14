package util

import (
	"encoding/json"
	"github.com/headzoo/surf"
	"io/ioutil"
	"mvdan.cc/xurls"
	"net/http"
	"strings"
)

type youtubeJSON struct {
	Kind  string  `json:"kind"`
	Items []Items `json:"items"`
}

type Items struct {
	Snippet Snippet `json:"snippet"`
}

type Snippet struct {
	ChannelId   string `json:"channelId"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func GetLinksFromYT() (linkList []string) {
	openYT := openYT("https://www.youtube.com/results?search_query=%22%3FcouponCode%3D%22&sp=EgQIARAB")
	var videoIDs []string

	if len(openYT) != 0 {
		for _, openLink := range openYT {
			videoID := strings.SplitAfter(openLink, "=")
			videoIDs = append(videoIDs, videoID[1])
		}
	}
	urlVideoIDs := strings.Join(videoIDs[:], ",")

	udemyLinks := scrapeYTAPI(urlVideoIDs)

	if len(udemyLinks) != 0 {
		linkList = udemyLinks
	}

	return
}

func openYT(url string) []string {
	var bow = surf.NewBrowser()
	bow.Open(url)
	var fURL string
	var videoLinks []string
	for _, link := range bow.Links() {
		if strings.Contains(link.URL.String(), "watch") && fURL != link.URL.String() {
			fURL = link.URL.String()
			videoLinks = append(videoLinks, link.URL.String())
		}
	}
	return videoLinks
}

func scrapeYTAPI(videoIDs string) (scrapedList []string) {

	var youtubeJson youtubeJSON
	client := &http.Client{}

	url := "https://www.googleapis.com/youtube/v3/videos?id=" + videoIDs + "&key=&part=snippet"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("User-Agent", "Golang_Spider_Bot/3.0")

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)

	json.Unmarshal([]byte(respBody), &youtubeJson)

	if len(youtubeJson.Items) != 0 {
		for _, desc := range youtubeJson.Items {
			urls := xurls.Relaxed().FindAllString(desc.Snippet.Description, -1)
			for _, udemylink := range urls {
				if strings.Contains(udemylink, "?couponCode=") {
					scrapedList = append(scrapedList, udemylink)
				}
			}
		}
	}
	return
}
