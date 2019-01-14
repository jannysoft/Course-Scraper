package util

import (
	"fmt"
	"strings"
	"encoding/json"
	"github.com/olivere/elastic"
	"context"
)

var elasticUrl = "http://35.204.247.102//elasticsearch/"

func AddToElasticSearch(course Course) {

	client := GetElasticCon(elasticUrl)

	putCourse, err := client.Index().Index("courses").Type("course").Id(course.ID).BodyJson(course).Do(context.Background())

	if err != nil {
		// Handle error
		panic(err)
	} else {
		fmt.Printf("Indexed course %s to index %s, type %s\n", putCourse.Id, putCourse.Index, putCourse.Type)
	}
}

func CheckCourseExists(id string) (course Course, error string) {

	client := GetElasticCon(elasticUrl)

	getCourse, err := client.Get().Index("courses").Type("course").Id(id).Do(context.Background())

	if err != nil {
		if strings.Contains(err.Error(), "404") {
			error = "404"
		}
	} else if getCourse.Found == true {
		json.Unmarshal(*getCourse.Source, &course)
	}
	return
}

func GetElasticCon(url string) *elastic.Client {
	client, err := elastic.NewSimpleClient(elastic.SetSniff(false), elastic.SetURL(url), elastic.SetBasicAuth("user", ""))

	if err != nil {
		panic(err)
	}

	return client
}
