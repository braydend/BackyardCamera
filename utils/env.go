package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func loadEnv(vars []string) (env map[string]string) {
	env = make(map[string]string)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .utils file")
	}

	for _, v := range vars {
		env[v] = os.Getenv(v)
	}

	return env
}


type env struct {
	Bucket string
	Region string
}

func GetEnv() *env {
	vars := loadEnv([]string{"BUCKET_NAME", "AWS_REGION"})
	bucket := vars["BUCKET_NAME"]
	region := vars["AWS_REGION"]

	return &env{bucket, region}
}