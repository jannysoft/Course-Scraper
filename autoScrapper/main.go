package main

import (
	"fmt"
	"scrapper/util"
	"math/rand"
	"time"
)

func main() {
	// Execute Main Loop
	for {
		//Create Global Scrape List
		var globalScrapeList []string
		fmt.Println(
			"====================================================================\n" +
				"A new scrape process has started, lets gather some links\n" +
				"====================================================================")

		// Get and add links from Discount Global to Global Scrape List
		fmt.Println("---- Gathering links from Discount Global ----")
		dgLinks := util.GetLinksFromDG()
		for _, dgLink := range dgLinks {
			globalScrapeList = append(globalScrapeList, dgLink)
		}
		fmt.Println("---- Adding", len(dgLinks), "links from Discount Global to Global Scrape List ----")
		// Get and add links from Learn Viral to Global Scrape List
		fmt.Println("---- Gathering links from Learn Viral ----")
		lvLinks := util.GetLinksFromLV()
		for _, lvLink := range lvLinks {
			globalScrapeList = append(globalScrapeList, lvLink)
		}
		fmt.Println("---- Adding", len(lvLinks), "links from Learn Viral to Global Scrape List ----")
		// Get and add links from Google 1 hour result search to Global Scrape List
		fmt.Println("---- Gathering links from Google Search ----")
		googleLinks := util.GetLinksFromGoogle()
		for _, googleLink := range googleLinks {
			globalScrapeList = append(globalScrapeList, googleLink)
		}
		fmt.Println("---- Adding", len(googleLinks), "links from Google Search to Global Scrape List ----")

		// Get and add links from YouTube 1 hour result search to Global Scrape List
		fmt.Println("---- Gathering links from YouTube Search ----")
		youtubeLinks := util.GetLinksFromYT()
		for _, youtubeLink := range youtubeLinks {
			globalScrapeList = append(globalScrapeList, youtubeLink)
		}
		fmt.Println("---- Adding", len(youtubeLinks), "links from YouTube Search to Global Scrape List ----")

		// Delete duplicate Links from Global Scrape List
		uniqueLinks := util.DeleteDuplicateLinks(globalScrapeList)
		fmt.Println("---- We have gathered", len(uniqueLinks), "links in total, it's time to scrape Udemy ----\n====================================================================")

		//Get Links From DB
		dbLinks := util.GetAllLinksFromDB()

		//Start Udemy Scrape Process
		for _, uniqueLink := range uniqueLinks {
			// Return bool when comparing links
			linkExistInDB := util.CompareLinks(dbLinks, uniqueLink)
			//If link does not exist in couchBase scrape it
			if linkExistInDB == false {
				fmt.Println("Currently Scrapping " + uniqueLink)

				course, err := util.ConstructUdemyCourse(uniqueLink)
				//Check for any errors
				if err !=nil {
					fmt.Println(err)
					udemyLink := util.AppendUdemyUrl(uniqueLink, err.Error())
					util.AppendLinkToDB(udemyLink)
				} else {
					// check if course with the same ID exists in elasticSearch
					checkCourseExists, errString := util.CheckCourseExists(course.ID)

					if errString == "404" {
						util.AddToElasticSearch(course)
						if course.CurrentPrice == "0" {
							hashTags := util.ScrapeHashTags()
							util.SocialSubmit(course, hashTags)
							fmt.Println("Course is discounted to FREE .. Should be added to all social accounts ! With extra hashtags: " + hashTags)
						}
						udemyLink := util.AppendUdemyUrl(uniqueLink, "added")
						util.AppendLinkToDB(udemyLink)
						fmt.Println("We Have succesfully added a course to DB !! Yay")

					} else if errString != "404" && checkCourseExists.CouponCode != course.CouponCode {
						util.AddToElasticSearch(course)
						if course.CurrentPrice == "0" {
							hashTags := util.ScrapeHashTags()
							util.SocialSubmit(course, hashTags)
							fmt.Println("Course is discounted to FREE .. Should be added to all social accounts ! With extra hashtags: " + hashTags)
						}
						udemyLink := util.AppendUdemyUrl(uniqueLink, "added")
						util.AppendLinkToDB(udemyLink)
						fmt.Println("We Have succesfully added a course to DB !! Yay")

					} else if errString != "404" && checkCourseExists.CouponCode == course.CouponCode {
						//fmt.Println("Course with the same coupon exists in database")
						udemyLink := util.AppendUdemyUrl(uniqueLink, "exists")
						util.AppendLinkToDB(udemyLink)
						fmt.Println("This course has not been added to DB.. Perhaps it's already there !")
					}
				}


				//Bot in standby Mode for about a minute
				randomD := rand.Intn(300-50) + 50
				fmt.Println("Standby for:", randomD, "seconds stay tooned !\n-----------------------------------------------")
				time.Sleep(time.Duration(randomD) * time.Second)
			}
		}

		// Shutdown the bot for about an hour !
		randomD := rand.Intn(60-30) + 30
		fmt.Println("All the websites and links has been processed accordingly! shutting down for approximatly: ", randomD, "minutes !")
		time.Sleep(time.Duration(randomD) * time.Minute)
	}
}