package s3

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"io"
	"net/url"
	"os"
	"time"
)

func getS3Client(region string) *s3.S3 {
	// This might get picked up from .utils
	az := aws.NewConfig().WithRegion(region)
	s3Session := session.Must(session.NewSession())
	return s3.New(s3Session, az)
}

func upload(body io.ReadSeeker, bucket, key, region string, timeout time.Duration) {
	s3Client := getS3Client(region)

	// Create a context with a timeout that will abort the upload if it takes
	// more than the passed in timeout.
	ctx := context.Background()
	var cancelFn func()
	if timeout > 0 {
		ctx, cancelFn = context.WithTimeout(ctx, timeout)
	}
	// Ensure the context is canceled to prevent leaking.
	// See context package for more information, https://golang.org/pkg/context/
	if cancelFn != nil {
		defer cancelFn()
	}

	// Uploads the object to S3. The Context will interrupt the request if the
	// timeout expires.
	_, err := s3Client.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   body,
	})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == request.CanceledErrorCode {
			// If the SDK can determine the request or retry delay was canceled
			// by a context the CanceledErrorCode error code will be returned.
			fmt.Fprintf(os.Stderr, "upload canceled due to timeout, %v\n", err)
		} else {
			fmt.Fprintf(os.Stderr, "failed to upload object, %v\n", err)
		}
		os.Exit(1)
	}

	fmt.Printf("successfully uploaded file to %s/%s\n", bucket, key)
}

func ListObjectsInBucket(bucket string) *s3.ListObjectsV2Output {
	client := getS3Client("ap-southeast-2")

	res, err := client.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: &bucket})

	if err != nil {
		panic(err)
	}

	return res
}

func UploadPhoto(filename, bucket, region string, timeout time.Duration) {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	upload(file, bucket, filename, region, timeout)
}

type ObjectWithUrl struct {
	Object *s3.Object
	Url string
}


func MapUrlsToObjects(objects []*s3.Object, urlRoot string) (objectsWithUrl []ObjectWithUrl) {
	for _, object := range objects {

		objectsWithUrl = append(objectsWithUrl, ObjectWithUrl{object, fmt.Sprintf("%s/%s", urlRoot, url.QueryEscape(*object.Key))})
	}

	return objectsWithUrl
}