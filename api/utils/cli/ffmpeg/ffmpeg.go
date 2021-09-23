package ffmpeg

import (
	"errors"
	"fmt"
	"github.com/p12s/okko-video-converter/api/common"
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

	arguments := fmt.Sprintf("-i %s -ss 00:00:01.000 -vframes 1 -filter:v scale='133:-1' %s", filePath, prevImageFilePath)

	_, err := cli.RunCli("ffmpeg", arguments)
	if err != nil {
		return "", err
	}

	return prevImageFilePath, nil
}

func ConvertVideo(path string, target common.FileType) error {
	if target == common.UNKNOWN {
		return errors.New(fmt.Sprintf("unknown file ext: %s\n", path))
	}
	if common.GetFileType(path) == target {
		fmt.Println("same type:", common.GetFileExt(target))
		return nil
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return errors.New("user directory or file does not exists, or somthing wrong with it")
	}

	targetFilePath := fs.GetPathWithoutExt(path) + "." + common.GetFileExt(target)
	arguments := fmt.Sprintf("-i %s -preset slow %s", path, targetFilePath)

	fmt.Println("run ConvertVideo:")

	_, err := cli.RunCli("ffmpeg", arguments)
	return err
}
