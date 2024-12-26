package utils

import (
	"io"
	"net/http"
	"os"
	"time"
)

const (
	IMG_URL      = "https://picsum.photos/1200"
	REFRESH_TIME = 10 * time.Second
)

func downloadNewImage(filePath string) error {
	response, err := http.Get(IMG_URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}

func GetImage() (string, error) {
	filePath, err := GetFilePath("img.png")
	if err != nil {
		return "", err
	}

	fileInfo, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		downloadNewImage(filePath)
		return filePath, nil
	} else if err != nil {
		return "", err
	}

	fileAge := time.Since(fileInfo.ModTime())
	if fileAge > REFRESH_TIME {
		downloadNewImage(filePath)
	}

	return filePath, nil
}
