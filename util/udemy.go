package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func ConstructUdemyCourse(url string) (Course, error) {

	var course Course
	// Open the submitted to standard out url
	response, err := urlRequest(url)
	if err != nil {
		return course, err
	}
	// Scrape the returned response for Course.ID
	scrapedComponents, err := scrapeComponents(response, url)
	if err != nil {
		return course, err
	}
	// Assign scraped components to course
	course = scrapedComponents

	return course, nil
}

func scrapeComponents(response *http.Response, url string) (Course, error) {

	var course Course
	defer response.Body.Close()
	// Open the returned response.Body with goquery and assign it to var doc
	doc, _ := goquery.NewDocumentFromReader(response.Body)
	if course.ID = doc.Find("body").AttrOr("data-clp-course-id", ""); course.ID == "" {
		return course, errors.New("Error: No Course ID Found !")
	}
	// Get the instructors for this course and append them to instructors string array
	selection := doc.Find("a.instructor-links__link")
	var instructors []string
	for i := range selection.Nodes {
		item := selection.Eq(i)
		scrapedInstructor := []byte(item.Text())
		re := regexp.MustCompile("  +")
		instructor := re.ReplaceAll(bytes.TrimSpace(scrapedInstructor), []byte(" "))
		instructors = append(instructors, string(instructor))
	}
	// Check if the course is created by banned by us instructor if it is return error
	isBannedInstructor := checkIfBannedInstructors(instructors)
	if isBannedInstructor == true {
		return course, errors.New("Error: This course is created by banned Instructor")
	}
	// Isolate the coupon code from the url
	coupon := strings.SplitAfter(url, "couponCode=")
	coupon = strings.Split(coupon[1], "&")
	// Construct the coupon/price url
	apiUrl := "https://www.udemy.com/api-2.0/course-landing-components/" + course.ID + "/me/?couponCode=" + coupon[0] + "&components=redeem_coupon,purchase"
	// Open the constructed api URL
	apiResponse, err := urlRequest(apiUrl)
	if err != nil {
		return course, err
	}
	// Unmarshal the json from apiResponse.Body to apiData var
	var apiData ApiData
	defer apiResponse.Body.Close()
	respBody, _ := ioutil.ReadAll(apiResponse.Body)
	json.Unmarshal([]byte(respBody), &apiData)
	// Check if the coupon code is valid
	if apiData.RedeemCoupon.Error != "" && apiData.RedeemCoupon.IsApplied == false {
		return course, errors.New("Error: " + apiData.RedeemCoupon.Error)
	}
	// Fill up the coupon data into course
	course.ValidCoupon = true
	course.CouponCode = apiData.RedeemCoupon.Code
	// Fill up the price data into course
	course.CurrentPrice = fmt.Sprintf("%v", apiData.Purchase.Discount.Price.Amount)
	course.OriginalPrice = fmt.Sprintf("%v", apiData.Purchase.Discount.ListPrice.Amount)
	course.DiscountRate = fmt.Sprintf("%v", apiData.Purchase.Discount.DiscountPercent)
	// Construct the url for the rest of the api Components
	apiUrl = "https://www.udemy.com/api-2.0/courses/" + course.ID +
		"?fields[course]=title,headline,primary_category,primary_subcategory,url,image_304x171,image_480x270,avg_rating,num_subscribers,locale"
	// Open the constructed api URL
	apiResponse, err = urlRequest(apiUrl)
	if err != nil {
		return course, err
	}
	// Unmarshal the json from apiResponse.Body to apiComponents var
	var apiComponents CourseComponents
	defer apiResponse.Body.Close()
	respBody, _ = ioutil.ReadAll(apiResponse.Body)
	json.Unmarshal([]byte(respBody), &apiComponents)
	// Fill up the data from apiComponents into course
	course.Title = apiComponents.Title
	course.Description = apiComponents.Headline
	course.Category = apiComponents.SubCategory.Title
	course.RemoteUrl = "https://www.udemy.com" + apiComponents.URL
	// course.PrettyURL and course.Slug
	reg, _ := regexp.Compile("[^A-Za-z0-9]+")
	prettyurl := reg.ReplaceAllString(course.Title, "-")
	prettyurl = strings.ToLower(strings.Trim(prettyurl, "-"))
	if len(prettyurl) == 0 {
		prettyurl = doc.Find("base").AttrOr("href", "")
		course.PrettyURL = reg.ReplaceAllString(prettyurl, "")
	} else {
		course.PrettyURL = prettyurl
	}
	catSlug := reg.ReplaceAllString(course.Category, "-")
	catSlug = strings.ToLower(strings.Trim(catSlug, "-"))
	course.CatSlug = catSlug

	course.Slug = "/" + course.ID + "/" + course.PrettyURL + ""

	course.PostIMG = apiComponents.PostIMG
	course.FrontIMG = apiComponents.FrontIMG
	course.Rating = fmt.Sprintf("%v", apiComponents.Rating)
	course.Enrolled = fmt.Sprintf("%v", apiComponents.Subscribers)
	course.Language = apiComponents.Language.EnglishTitle
	course.Timestamp = time.Now().Unix()
	//GET course.WillLearn
	pulledValues := doc.Find(".what-you-get__text").Map(func(i int, s *goquery.Selection) string {
		return s.Text()
	})
	course.WillLearn = pulledValues
	return course, nil
}

func checkIfBannedInstructors(instructors []string) bool {

	var isBannedInstructors []bool

	for _, instructor := range instructors {
		switch instructor {
		case "ALL TECH LEARN":
			isBannedInstructors = append(isBannedInstructors, true)
		case "KUNCHAM Software Solutions Pvt Ltd":
			isBannedInstructors = append(isBannedInstructors, true)
		case "Zeuxin Solutions":
			isBannedInstructors = append(isBannedInstructors, true)
		case "Wilskill Technologies":
			isBannedInstructors = append(isBannedInstructors, true)
		case "Emposes Solutions":
			isBannedInstructors = append(isBannedInstructors, true)
		case "Techies Tube":
			isBannedInstructors = append(isBannedInstructors, true)
		default:
			isBannedInstructors = append(isBannedInstructors, false)
		}
	}
	for _, isBannedInstructor := range isBannedInstructors {
		if isBannedInstructor == true {
			return true
		}
	}
	return false
}

func urlRequest(url string) (*http.Response, error) {
	browser := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")
	response, err := browser.Do(request)

	if err != nil {
		return nil, err
	}
	return response, nil
}
