package util

import (
	"github.com/headzoo/surf"
	"strconv"
	"strings"
)

func GetLinksFromDG() (linkList []string) {
	linkList = openSite("http://udemycoupon.discountsglobal.com/page/")
	return
}

func GetLinksFromLV() (linkList []string) {
	linkList = openSite("https://udemycoupon.learnviral.com/page/")
	return
}

func openSite(siteUrl string) (scrapedList []string) {
	bow := surf.NewBrowser()
	for i := 1; i <= 3; i++ {

		s := strconv.Itoa(i)
		url := siteUrl + s + "/"
		bow.Open(url)
		sel := bow.Find("div.link-holder")

		for i := range sel.Nodes {
			item := sel.Eq(i)
			linkTag := item.Find("a")
			scrapedLink, _ := linkTag.Attr("href")
			if strings.Contains(scrapedLink, "/?couponCode=") {
				scrapedList = append(scrapedList, scrapedLink)
			}
		}
	}
	return
}
