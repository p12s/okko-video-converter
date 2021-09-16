package handler

import (
	"encoding/json"
	"fmt"
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

func (h *Handler) updateResizeOptions(c *gin.Context) {
	userCode, err := getUserCode(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	var options common.UserData
	if err := c.BindJSON(&options); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	optionsAsJson, err := json.Marshal(options)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.services.ResizeOptions.UpdateOrCreate(common.ResizeOptions{
		Options: string(optionsAsJson),
		Status:  common.CREATED,
	}, userCode)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	go h.runResize(userCode)

	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "OK",
	})
}

// getResizeProgress
func (h *Handler) getResizeProgress(c *gin.Context) {
	userCode, err := getUserCode(c)
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	resizeOptions, err := h.services.ResizeOptions.Get(userCode)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	progress := 100
	if resizeOptions.TotalCount > 0 {
		progress = (resizeOptions.Current * 100) / resizeOptions.TotalCount
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"progress": progress,
	})
}

func (h *Handler) runResize(userCode string) {
	fmt.Println("RUN Gorutine resize")

	// _, err := fs.IsFilesExists(userCode)
	// if err != nil {
	// 	saveErr := h.services.ResizeOptions.SaveError(userCode, err.Error())
	// 	fmt.Println("runResize.IsFilesExists error: ", err.Error(), "ResizeOptions.SaveError error: ", saveErr.Error())
	// 	return
	// }

	// options, err := h.services.ResizeOptions.Get(userCode)
	// if err != nil {
	// 	saveErr := h.services.ResizeOptions.SaveError(userCode, err.Error())
	// 	fmt.Println("runResize.ResizeOptions.Get error: ", err.Error(), "ResizeOptions.SaveError error: ", saveErr.Error())
	// 	return
	// }

	// files, err := h.services.File.GetAll(userCode)
	// if err != nil || len(files) == 0 {
	// 	saveErr := h.services.ResizeOptions.SaveError(userCode, err.Error())
	// 	fmt.Println("runResize.File.GetAll error: ", err.Error(), "ResizeOptions.SaveError error: ", saveErr.Error())
	// 	return
	// }

	// var userData common.UserData
	// if err := json.Unmarshal([]byte(options.Options), &userData); err != nil {
	// 	saveErr := h.services.ResizeOptions.SaveError(userCode, err.Error())
	// 	fmt.Println("runResize.json.Unmarshal error: ", err.Error(), "ResizeOptions.SaveError error: ", saveErr.Error())
	// 	return
	// }

	// widthList := calc.GenerateWidthList(userData)

	// // удалим папку польз. с вложениями (если она есть) и создадим заново
	// destFolder := os.Getenv("FILE_DIR") + userCode + "/result"
	// err = os.RemoveAll(destFolder)
	// if err != nil {
	// 	saveErr := h.services.ResizeOptions.SaveError(userCode, err.Error())
	// 	fmt.Println("runResize.os.RemoveAll error: ", err.Error(), "ResizeOptions.SaveError error: ", saveErr.Error())
	// 	return
	// }
	// err = os.Mkdir(destFolder, 0755)
	// if err != nil {
	// 	saveErr := h.services.ResizeOptions.SaveError(userCode, err.Error())
	// 	fmt.Println("runResize.os.Mkdir error: ", err.Error(), "ResizeOptions.SaveError error: ", saveErr.Error())
	// 	return
	// }

	// file := files[0]
	// resizedImages := []common.ResizedImg{}

	// currentFile := 0
	// for _, item := range widthList.List {
	// 	for _, field := range item.SizeFields {

	// 		targetWidth := uint(float64(item.Width) * field.Coef)
	// 		if targetWidth < uint(file.Width) {
	// 			imgs, err := fs.TryResizeImage(file, userCode, targetWidth, field.Coef,
	// 				userData.IsCompress, userData.IsAddWebp)
	// 			if err != nil {
	// 				saveErr := h.services.ResizeOptions.SaveError(userCode, err.Error())
	// 				fmt.Println("runResize.fs.TryResizeImage error: ", err.Error(), "ResizeOptions.SaveError error: ", saveErr.Error())
	// 				return
	// 			}
	// 			resizedImages = append(resizedImages, imgs...)
	// 		}

	// 		currentFile += 1
	// 		err = h.services.ResizeOptions.UpdateTotalAndCurrent(userCode, widthList.Total, currentFile)
	// 		if err != nil {
	// 			saveErr := h.services.ResizeOptions.SaveError(userCode, err.Error())
	// 			fmt.Println("runResize.ResizeOptions.UpdateTotalAndCurrent error: ", err.Error(), "ResizeOptions.SaveError error: ", saveErr.Error())
	// 			return
	// 		}
	// 	}
	// }

	// err = fs.Archive(userCode, resizedImages)
	// if err != nil {
	// 	saveErr := h.services.ResizeOptions.SaveError(userCode, err.Error())
	// 	fmt.Println("runResize.fs.Archive error: ", err.Error(), "ResizeOptions.SaveError error: ", saveErr.Error())
	// 	return
	// }

	// err = h.services.ResizeOptions.UpdateFinishTime(userCode)
	// if err != nil {
	// 	saveErr := h.services.ResizeOptions.SaveError(userCode, err.Error())
	// 	fmt.Println("runResize.ResizeOptions.UpdateFinishTime error: ", err.Error(), "ResizeOptions.SaveError error: ", saveErr.Error())
	// 	return
	// }

	// fmt.Println("Resize ended OK")
	// return
}

func (h *Handler) resize(c *gin.Context) {
	// token := c.Param("token")
	// if token == "" {
	// 	NewErrorResponse(c, http.StatusBadRequest, "invalid token param")
	// 	return
	// }
	// userUUID, err := h.services.ParseToken(token)
	// if userUUID == uuid.Nil {
	// 	NewErrorResponse(c, http.StatusBadRequest, "invalid userCode param")
	// 	return
	// }
	// userCode := userUUID.String()

	// isFilesExists, err := fs.IsFilesExists(userCode)
	// if !isFilesExists || err != nil {
	// 	NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// // достаю из бд настройки
	// resizeOptions, err := h.services.ResizeOptions.Get(userCode)
	// if err != nil {
	// 	NewErrorResponse(c, http.StatusUnauthorized, err.Error())
	// 	return
	// }

	// var options common.UserData
	// if err := json.Unmarshal([]byte(resizeOptions.Options), &options); err != nil {
	// 	NewErrorResponse(c, http.StatusBadRequest, err.Error())
	// 	return
	// }

	// widthList := calc.GenerateWidthList(options)

	// resizeProcess := common.ResizeProcess{
	// 	Total:   widthList.Total,
	// 	Current: 0,
	// }

	// files, err := h.services.File.GetAll(userCode)
	// if err != nil || len(files) == 0 {
	// 	NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// // удалим эту папку с вложениями (если она есть) и создадим заново
	// destFolder := os.Getenv("FILE_DIR") + userCode + "/result"
	// err = os.RemoveAll(destFolder)
	// if err != nil {
	// 	NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }
	// err = os.Mkdir(destFolder, 0755)
	// if err != nil {
	// 	NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// file := files[0]
	// resizedImages := []common.ResizedImg{}

	// c.Stream(func(w io.Writer) bool {

	// 	for _, item := range widthList.List {
	// 		for _, field := range item.SizeFields {

	// 			targetWidth := uint(float64(item.Width) * field.Coef)

	// 			imgs, err := fs.TryResizeImage(file, userCode, targetWidth, field.Coef, options.IsCompress, options.IsAddWebp)
	// 			if err != nil {
	// 				NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 				return false
	// 			}
	// 			resizedImages = append(resizedImages, imgs...)

	// 			resizeProcess.Current += 1
	// 			percent := resizeProcess.Current * 100 / resizeProcess.Total
	// 			fmt.Printf("%d%%, %d шт.\n", percent, len(resizedImages))

	// 			process, err := json.Marshal(resizeProcess)
	// 			if err != nil {
	// 				NewErrorResponse(c, http.StatusBadRequest, err.Error())
	// 				return false
	// 			}
	// 			c.SSEvent("message", string(process))
	// 		}

	// 	}

	// 	err = fs.Archive(userCode, resizedImages)
	// 	if err != nil {
	// 		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
	// 		return false
	// 	}

	// 	fmt.Println("stream - STOP")
	// 	// save into BD status, time

	// 	resizeProcess.Current = resizeProcess.Total
	// 	process, err := json.Marshal(resizeProcess)
	// 	if err != nil {
	// 		NewErrorResponse(c, http.StatusBadRequest, err.Error())
	// 		return false
	// 	}
	// 	c.SSEvent("message", string(process))

	// 	return false
	// })
}

