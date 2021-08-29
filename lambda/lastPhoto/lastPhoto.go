package main

import (
	"context"
	"encoding/json"
	"fmt"
	"gardenPhotosS3/s3"
	"github.com/aws/aws-lambda-go/lambda"
	"net/url"
	"sort"
)

type Response struct {
	Filename string `json:"filename"`
	URL string `json:"url"`
}

func HandleRequest(ctx context.Context) (string, error) {
	objects := s3.ListObjectsInBucket("backyard-photos").Contents

	sort.Slice(objects, func(i, j int) bool {
		return objects[i].LastModified.UnixNano() > objects[j].LastModified.UnixNano()
	})

	newestObject := objects[0]
	s3Url := fmt.Sprintf("https://backyard-photos.s3.ap-southeast-2.amazonaws.com/%s", url.QueryEscape(*newestObject.Key))

	response, err := json.Marshal(Response{*newestObject.Key, s3Url})

	if err != nil {
		panic(err)
	}

	return string(response), nil
}

func main() {
	lambda.Start(HandleRequest)
}