package main

import (
	"fmt"
	"gardenPhotosS3/fswebcam"
	"gardenPhotosS3/s3"
	"gardenPhotosS3/utils"
	"net/http"
	"time"
)

func handleTakePhoto(w http.ResponseWriter, req *http.Request) {
	filename := time.Now().Format(time.RFC3339) + ".jpg"
	timeout, err := time.ParseDuration("5m")

	if err != nil {
		panic(err)
	}

	env := utils.GetEnv()

	fswebcam.TakePhoto(filename)
	s3.UploadPhoto(filename, env.Bucket, env.Region, timeout)
	utils.DeleteFile(filename)
	_, err = w.Write([]byte(fmt.Sprintf("Photo %s was taken successfully!", filename)))

	if err != nil {
		panic(err)
	}
}

func main() {

	http.HandleFunc("/takePhoto", handleTakePhoto)

	err := http.ListenAndServe(":8090", nil)

	if err != nil {
		panic(err)
	}
}