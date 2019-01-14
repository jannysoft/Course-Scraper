package util

import "time"

var udemyLink UdemyLink

func DeleteDuplicateLinks(scrapedLinks []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, link := range scrapedLinks {
		if _, value := keys[link]; !value {
			keys[link] = true
			list = append(list, link)
		}
	}
	return list
}

func CompareLinks(dbLinks []string, uniqueLink string) bool {
	for _, dbLink := range dbLinks {
		if dbLink == uniqueLink {
			return true
		}
	}
	return false
}

func AppendUdemyUrl(link string, message string) UdemyLink {
	udemyLink.Link = link
	udemyLink.Message = message
	udemyLink.Time = time.Now().Unix()
	return udemyLink
}

