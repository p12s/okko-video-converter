package ffmpeg

import (
	"errors"
	"fmt"
	"github.com/p12s/okko-video-converter/api/utils/cli"
	"github.com/p12s/okko-video-converter/api/utils/fs"
	"os"
)

func GenerateVideoPreview(userCode, fileName string) (string, error) {
	filePath := os.Getenv("FILE_DIR") + userCode + "/" + fileName
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", errors.New("user directory or file does not exists, or somthing wrong with it")
	}

	prevImageName := fs.FileNameWithoutExt(fileName) + ".jpg"
	prevImageFilePath := os.Getenv("FILE_DIR") + userCode + "/" + prevImageName
	fmt.Println("prevImageFilePath:", prevImageFilePath)

	arguments := fmt.Sprintf("-i %s -ss 00:00:01.000 -vframes 1 -filter:v scale='133:-1' %s", filePath, prevImageFilePath)

	_, err := cli.RunCli("ffmpeg", arguments)
	if err != nil {
		return "", err
	}

	return prevImageFilePath, nil
}
