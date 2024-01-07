package go_utils

import (
	zip33 "archive/zip"
	"log"
)

func RmZipFile(zipFile string, rmFs ...string) {
	jarWriter, err := zip33.OpenReader(zipFile)
	if err != nil {
		log.Println(err)
		return
	}
	defer jarWriter.Close()
	for _, file := range jarWriter.File {
		for _, x := range rmFs {
			if file.Name == x {
				file.Remove()

			}
		}
	}
}
