package utils

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func LoadEnv(vars []string) (env map[string]string) {
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
