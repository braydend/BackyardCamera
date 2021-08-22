package main

import (
	"flag"
	"gardenPhotosS3/fswebcam"
	"gardenPhotosS3/s3"
	"gardenPhotosS3/utils"
	"os"
	"time"
)

type flags struct {
	timeout time.Duration
}

type env struct {
	bucket string
	region string
}

func deleteFile(filename string) {
	err := os.Remove(filename)

	if err != nil {
		panic(err)
	}
}

func takePhoto(filename string) {
	fswebcam.TakePhoto(filename)
}

func uploadPhoto(filename, bucket, region string, timeout time.Duration) {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	s3.Upload(file, bucket, "testphoto.jpg", region, timeout)
}

func parseFlags() flags {
	var timeout time.Duration

	flag.DurationVar(&timeout, "d", 0, "Upload timeout.")
	flag.Parse()

	return flags{timeout}
}

func getEnv() env {
	vars := utils.LoadEnv([]string{"BUCKET_NAME", "AWS_REGION"})
	bucket := vars["BUCKET_NAME"]
	region := vars["AWS_REGION"]

	return env{bucket, region}
}

func main() {
	filename := time.Now().Format(time.RFC3339) + ".jpg"

	env := getEnv()
	flags := parseFlags()

	takePhoto(filename)
	uploadPhoto(filename, env.bucket, env.region, flags.timeout)
	deleteFile(filename)
}
