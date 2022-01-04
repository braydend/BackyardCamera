package main

import (
	"encoding/json"
	"fmt"
	"getPhotos/s3"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"sort"
	"strconv"
	"time"
)

const defaultLength = 10

func buildDatePrefix(date time.Time) string {
	month := date.Month()
	year := date.Year()

	return fmt.Sprintf("%d-%02d", year, month)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	objectPrefix := buildDatePrefix(time.Now())
	objects, err := s3.ListObjectsInBucket("backyard-photos", objectPrefix)

	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{}, err
	}

	limit, err := strconv.Atoi(request.QueryStringParameters["limit"])

	if err != nil {
		limit = defaultLength
	}

	if limit > len(objects.Contents) {
		limit = len(objects.Contents)
	}

	sort.Slice(objects.Contents, func(i, j int) bool {
		return objects.Contents[i].LastModified.UnixNano() > objects.Contents[j].LastModified.UnixNano()
	})

	objectsWithUrl := s3.MapUrlsToObjects(objects.Contents[0:limit], "https://backyard-photos.s3.ap-southeast-2.amazonaws.com")
	response, err := json.Marshal(objectsWithUrl)

	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{}, err
	}

	return events.APIGatewayProxyResponse{
		Body:       string(response),
		StatusCode: 200,
		Headers: map[string]string{
			"Access-Control-Allow-Headers": "Content-Type",
			"Access-Control-Allow-Origin":  "*",   // Allow from anywhere
			"Access-Control-Allow-Methods": "GET", // Allow only GET request
		},
	}, nil
}

func main() {
	lambda.Start(handler)
}
