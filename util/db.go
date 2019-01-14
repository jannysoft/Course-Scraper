package util

import (
	"encoding/json"
	"fmt"
	"github.com/couchbase/gocb"
	"context"
)

var bucket *gocb.Bucket

func init() {
	//Connect to Couchbase
	cluster, _ := gocb.Connect("couchbase://192.168.1.85")
	bucket, _ = cluster.OpenBucket("UdemyCourses", "")
}

func GetAllLinksFromDB() (dbLinks []string) {
	//var n1qlParams []interface{}
	//n1qlParams = append(n1qlParams, "udemyUrls")
	var dbData DBData
	dbQuery := gocb.NewN1qlQuery("SELECT links FROM UdemyCourses")
	dbResult, _ := bucket.ExecuteN1qlQuery(dbQuery, nil)
	json.Unmarshal(dbResult.NextBytes(), &dbData)
	for _, dbLink := range dbData.Links {
		dbLinks = append(dbLinks, dbLink.Link)
	}
	dbResult.Close()
	return
}

func DeleteCourse(id string) {
	client := GetElasticCon(elasticUrl)
	client.Delete().Index("courses").Type("course").Id(id).Do(context.Background())
}

func AppendLinkToDB(link UdemyLink) {
	var n1qlParams []interface{}
	n1qlParams = append(n1qlParams, "udemyUrls")
	n1qlParams = append(n1qlParams, link.Link)
	n1qlParams = append(n1qlParams, link.Message)
	n1qlParams = append(n1qlParams, link.Time)
	dbQuery := gocb.NewN1qlQuery("update UdemyCourses USE KEYS $1 SET links = ARRAY_APPEND( links, {'link':$2, 'message':$3, 'time':$4})")
	dbResult, err := bucket.ExecuteN1qlQuery(dbQuery, n1qlParams)
	if err != nil {
		fmt.Println(err)
	}
	dbResult.Close()
}
