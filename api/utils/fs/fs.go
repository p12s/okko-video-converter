package fs

import (
	"os"
	"path/filepath"
)

func CreateDirIfNotExists(dirPath string) error {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return os.Mkdir(dirPath, 0755)
	}
	return nil
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

func ClearAllUserFiles(userCode string) error {
	userDir := os.Getenv("FILE_DIR") + userCode
	err := os.RemoveAll(userDir + ".zip")
	if err != nil {
		return err
	}
	return os.RemoveAll(userDir)
}

func FileNameWithoutExt(fileName string) string {
	return fileName[:len(fileName)-len(filepath.Ext(fileName))]
}

func GetPathWithoutExt(path string) string {
	if len(path) == 0 {
		return ""
	}
	extIndex := len(path) - 1
	for i := len(path) - 1; i >= 0; i-- {
		if path[i] == '.' {
			extIndex = i
			break
		}
	}
	return path[0:extIndex]
}
