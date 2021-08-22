package main

import (
	"gardenPhotosS3/fswebcam"
	"time"
)

func main() {
	filename := time.Now().Format(time.RFC3339) + ".jpg"

	fswebcam.TakePhoto(filename)
}
