package main

import (
	"encoding/json"
	"fmt"
	"getPhotos/s3"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"sort"
	"strconv"
)

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	objects, err := s3.ListObjectsInBucket("backyard-photos")

	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{}, err
	}

	limit, err := strconv.Atoi(request.QueryStringParameters["limit"])

	if err != nil {
		limit = 10
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
	}, nil
}

func main() {
	lambda.Start(handler)
}