package utils

import "os"

func DeleteFile(filename string) {
	err := os.Remove(filename)

	if err != nil {
		panic(err)
	}
}
