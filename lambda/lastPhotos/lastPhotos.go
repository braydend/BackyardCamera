package main

import (
	"encoding/json"
	"gardenPhotosS3/s3"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"sort"
	"strconv"
)

func HandleRequest(request events.APIGatewayProxyRequest) (string, error) {
	objects := s3.ListObjectsInBucket("backyard-photos").Contents

	limit, err := strconv.Atoi(request.QueryStringParameters["limit"])

	if err != nil {
		limit = 10
	}

	sort.Slice(objects, func(i, j int) bool {
		return objects[i].LastModified.UnixNano() > objects[j].LastModified.UnixNano()
	})

	objectsWithUrl := s3.MapUrlsToObjects(objects[0:limit], "https://backyard-photos.s3.ap-southeast-2.amazonaws.com")
	response, err := json.Marshal(objectsWithUrl)

	if err != nil {
		return "", err
	}

	return string(response), nil
}

func main() {
	lambda.Start(HandleRequest)
}