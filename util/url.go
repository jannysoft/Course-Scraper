package util

import (
	"strings"
	"github.com/headzoo/surf"
)

var prevUrl string

func CheckUrl(url string) (string, bool) {

	bow := surf.NewBrowser()

	openURL := bow.Open(url)

	if openURL != nil {
		//fmt.Println(openURL)

	} else {
		url := bow.Url().String()

		if prevUrl == url {
			//fmt.Println("prevURL do nothing")
		} else {
			//fmt.Println("+++ URL AFTER SURF +++")
			if strings.Contains(url, "https://www.udemy.com/courses/") {
				//fmt.Println(color.YellowString("Link To All Courses:"), url, ">>", "shutting down")
				return "", false
			} else if strings.Contains(url, "https://www.udemy.com/collection/") {
				//fmt.Println(color.YellowString("Link To Collection:"), url, ">>", "shutting down")
				return "", false
			} else if strings.Contains(url, "https://www.udemy.com/draft/") {
				//fmt.Println(color.YellowString("Link To Draft:"), url, ">>", "shutting down")
				return "", false
			} else if strings.Contains(url, "https://www.udemy.com/certificate/") {
				//fmt.Println(color.YellowString("Link To Certificate:"), url, ">>", "shutting down")
				return "", false
			} else if strings.Contains(url, "https://www.udemy.com/s/") {
				//fmt.Println(color.YellowString("Link To Certificate:"), url, ">>", "shutting down")
				return "", false
			} else if strings.Contains(url, "https://www.udemy.com/user/") {
				//fmt.Println(color.YellowString("Link To User:"), url, ">>", "shutting down")
				return "", false
			} else if strings.Contains(url, "https://www.udemy.com/blog/") {
				//fmt.Println(color.YellowString("Link To Blog:"), url, ">>", "shutting down")
				return "", false
			} else if strings.Contains(url, "couponCode=") {
				//fmt.Println(color.GreenString("UDEMY LINK FOUND:"), url, ">>", color.GreenString("LETS SCRAPE"))
				return url, true
			} else if strings.Contains(url, "https://www.dennisjsmith.com") {
				//fmt.Println(color.RedString("SHIT LINK"), url, ">>", "shutting down")
				return "", false
			} else {
				//fmt.Println(color.YellowString("Unknown Link:"), url, ">>", "shutting down")
				return "", false
			}
		}
		//fmt.Println(prevUrl)
		prevUrl = url
	}
	return "", false
}
