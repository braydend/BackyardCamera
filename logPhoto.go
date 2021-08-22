package main

import (
	"flag"
	"gardenPhotosS3/fswebcam"
	"gardenPhotosS3/s3"
	"gardenPhotosS3/utils"
	"time"
)

type flags struct {
	timeout time.Duration
}

func parseFlags() flags {
	var timeout time.Duration

	flag.DurationVar(&timeout, "d", 0, "Upload timeout.")
	flag.Parse()

	return flags{timeout}
}

func main() {
	filename := time.Now().Format(time.RFC3339) + ".jpg"

	env := utils.GetEnv()
	flags := parseFlags()

	fswebcam.TakePhoto(filename)
	s3.UploadPhoto(filename, env.Bucket, env.Region, flags.timeout)
	utils.DeleteFile(filename)
}
