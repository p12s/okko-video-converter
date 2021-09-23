package handler

import (
	"fmt"
	"github.com/p12s/okko-video-converter/api/pkg/broker"
	"github.com/p12s/okko-video-converter/api/utils/cli/ffmpeg"
	"github.com/p12s/okko-video-converter/api/utils/validator"
	"github.com/spf13/viper"
	"os"

	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/p12s/okko-video-converter/api/common"
	"github.com/p12s/okko-video-converter/api/utils/fs"
)

func (h *Handler) removeAll(c *gin.Context) {
	userCode, err := getUserCode(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	err = fs.ClearAllUserFiles(userCode)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "Can't clear user files: "+err.Error())
		return
	}

	err = h.services.File.DeleteAll(userCode)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "Can't remove from DB user files: "+err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "OK",
	})
}

func (h *Handler) upload(c *gin.Context) {
	userCode, err := getUserCode(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	userImgDir := os.Getenv("FILE_DIR") + userCode
	err = fs.CreateDirIfNotExists(userImgDir)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	var uploadedFiles []common.UploadedFile

	form, err := c.MultipartForm()
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "Form parse error: "+err.Error())
		return
	}

	err = fs.ReCreateUserDir(userCode)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	err = h.services.File.DeleteAll(userCode)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	extension := ""
	if val, ok := form.Value["extension"]; ok {
		if len(val) > 0 {
			extension = val[0]
		}
	}

	files := form.File["files"]
	for _, file := range files {

		item := common.UploadedFile{
			Name:         file.Filename,
			KiloByteSize: file.Size / 1000,
			Type:         common.GetFileType(file.Filename),
		}

		fileType := strings.Join(file.Header["Content-Type"], "")
		if !validator.IsFileTypeValid(fileType) {
			item.Error = "Wrong file type: " + fileType
			item.IsError = true
			uploadedFiles = append(uploadedFiles, item)
			continue
		}

		if !validator.IsFileSizeExceeded(item.KiloByteSize) {
			item.Error = fmt.Sprintf("File oversized! Allowed max size is %dMb", viper.GetInt64("maxFileSizeInMb"))
			item.IsError = true
			uploadedFiles = append(uploadedFiles, item)
			continue
		}

		filePath := os.Getenv("FILE_DIR") + userCode + "/" + file.Filename
		if err := c.SaveUploadedFile(file, filePath); err != nil {
			item.Error = "Upload file err: " + err.Error()
			item.IsError = true
			uploadedFiles = append(uploadedFiles, item)
			continue
		}
		item.Path = filePath

		prevPath, err := ffmpeg.GenerateVideoPreview(userCode, file.Filename)
		if err != nil {
			item.Error = "Generate video preview err: " + err.Error()
			item.IsError = true
			uploadedFiles = append(uploadedFiles, item)
			continue
		}
		item.PrevImage = prevPath

		uploadedFiles = append(uploadedFiles, item)
	}

	var correctImages []common.UploadedFile
	for _, v := range uploadedFiles {
		if !v.IsError {
			// ставим задачу в очередь
			err = broker.Create(h.broker.Producer, common.VideoConvertData{
				UserCode:     userCode,
				Path:         v.Path,
				TargetFormat: common.GetFileType(extension),
			})
			if err != nil {
				NewErrorResponse(c, http.StatusInternalServerError, err.Error())
				return
			}
			correctImages = append(correctImages, v)
		}
	}

	err = h.services.File.Create(correctImages, userCode)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"files": uploadedFiles,
	})
}

func (h *Handler) download(c *gin.Context) {
	userCode, err := getUserCode(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	if !fs.IsUserArchiveExists(userCode) {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+userCode+".zip")
	c.Header("Content-Type", "application/octet-stream")
	c.File(os.Getenv("FILE_DIR") + userCode + ".zip")
}

func (h *Handler) files(c *gin.Context) {
	userCode, err := getUserCode(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	items, err := h.services.File.GetAll(userCode)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, items)
}
