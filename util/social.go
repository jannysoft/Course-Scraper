package util

import (
	"github.com/MariaTerzieva/gotumblr"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/jzelinskie/geddit"
	"net/http"
)

var pinterestAPI = "https://api.pinterest.com/v1/pins/"

const (
	consumerKey    = ""
	consumerSecret = ""
	accessToken    = "90209278-"
	accessSecret   = ""
)

func SocialSubmit(course Course, hashtags string) {
	//facebook(course)
	pinterest(course, hashtags)
	reddit(course)
	//tumblr(course)
	twittr(course, hashtags)
}

//func facebook(course Course) {
//	res, _ := fb.Post("/373495176180733/feed", fb.Params{
//		"description":  course.Description,
//		"name":         course.Title,
//		"picture":      course.PostIMG,
//		"link":         "https://www.greatcourses.co/course/" + course.ID + "/" + course.PrettyURL + "/",
//		"access_token": "",
//	})
//}

func pinterest(course Course, hashtags string) {
	client := &http.Client{}

	req, _ := http.NewRequest("POST", pinterestAPI, nil)

	query := req.URL.Query()
	query.Add("access_token", "")
	query.Add("board", "602989906297568207")
	query.Add("note", "[UDEMY 100% OFF]"+course.Title+"-"+course.Description+"#udemy #udemycoupon #udemyfreecoupon #udemycouponcode "+ hashtags)
	query.Add("image_url", course.PostIMG)
	query.Add("link", "https://www.greatcourses.co/goto/"+course.ID)
	req.URL.RawQuery = query.Encode()

	client.Do(req)
}

func reddit(course Course) {
	sesh, err := geddit.NewLoginSession("", "", "testing user agent")

	if err != nil {
		panic(err)
	}

	sub := geddit.NewLinkSubmission("udemyfreebies", "[100% OFF] "+course.Title, "https://www.greatcourses.co/goto/"+course.ID, true, &geddit.Captcha{})

	sesh.Submit(sub)
}

func tumblr(course Course) {
	blogname := "greatcoursesco.tumblr.com"
	state := "published"

	client := gotumblr.NewTumblrRestClient("",
		"",
		"",
		"",
		"https://www.greatcourses.co",
		"http://api.tumblr.com")

	title := "[UDEMY 100% OFF] " + course.Title
	url := "https://www.greatcourses.co/goto/" + course.ID
	description := course.Description
	thumbnail := course.PostIMG
	excerpt := ""
	tags := "udemy, freecourse, education, free"
	client.CreateLink(
		blogname, map[string]string{"url": url, "title": title, "description": description, "thumbnail": thumbnail, "excerpt": excerpt, "tags": tags, "state": state})

}

func twittr(course Course, hashtags string) {
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	client := twitter.NewClient(httpClient)
	client.Statuses.Update("[100% OFF] "+course.Title+ " " + hashtags +" #udemy #coupon #free #udemycourse #education https://www.greatcourses.co/course/"+course.ID+"/"+course.PrettyURL+"/", nil)

}
