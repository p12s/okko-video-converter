package fs

import (
	"errors"
	"io"
	"os"
	"path/filepath"

	"github.com/p12s/okko-video-converter/api/common"
)

func CreateDirIfNotExists(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return os.Mkdir(dirPath, 0755)
	}
	return nil
}

func IsFilesExists(userCode string) (bool, error) {
	folderPath := os.Getenv("FILE_DIR") + userCode

	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		return false, errors.New("user directory does not exists, or somthing wrong with it")
	}

	directory, err := os.Open(folderPath)
	if err != nil {
		return false, errors.New("error while directory was open")
	}
	defer directory.Close()

	// read in ONLY one file
	_, err = directory.Readdir(1)
	// and if the file is EOF... well, the dir is empty.
	if err == io.EOF {
		return false, errors.New("user directory is empty")
	}

	return true, nil
}

func IsUserArchiveExists(userCode string) bool {
	if _, err := os.Stat(os.Getenv("FILE_DIR") + userCode + ".zip"); os.IsNotExist(err) {
		return false
	}
	return true
}

func ReCreateUserDir(userCode string) error {
	userDir := os.Getenv("FILE_DIR") + userCode
	err := os.RemoveAll(userDir)
	if err != nil {
		return err
	}
	err = os.Mkdir(userDir, 0755)
	if err != nil {
		return err
	}
	return nil
}

// Обработка одной картинки под один размер, сжатие, создание копии в webp - и возврат путей для архиватора
func TryResizeImage(
	file common.File, userCode string,
	width uint, coef float64,
	isShouldCompress, isAddWebp bool) ([]common.ResizedImg, error) {

	resizedImages := []common.ResizedImg{}
	return resizedImages, nil
}

// Обработка всех картинок пользователя в его директории (не смотрим в БД)
func TryResizeImages(userCode string, widthList common.WidthList, isShouldCompress bool) ([]common.ResizedImg, error) {

	resizedImages := []common.ResizedImg{}

	return resizedImages, nil
}

func ResizeImage(fileSrc, fileDest string, width, height uint, isShouldCompress bool) error {

	return nil
}

func Archive(userCode string, resizedImg []common.ResizedImg) error {

	return nil
}

func ClearAllUserFiles(userCode string) error {
	userDir := os.Getenv("FILE_DIR") + userCode
	err := os.RemoveAll(userDir + ".zip")
	if err != nil {
		return err
	}
	return os.RemoveAll(userDir)
}

func ClearFile(userCode, fileName string) error {
	return os.RemoveAll(os.Getenv("FILE_DIR") + userCode + "/" + fileName)
}

func FileNameWithoutExt(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}