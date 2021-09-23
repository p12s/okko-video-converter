package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/p12s/okko-video-converter/api/common"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) websocket(c *gin.Context) {
	ws, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("error get connection")
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	defer ws.Close()

	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	var data struct {
		Token string `json:"token"`
	}
	err = ws.ReadJSON(&data)
	if err != nil || data.Token == "" {
		fmt.Println("error read token")
		return
	}

	userCode, err := h.services.User.ParseToken(data.Token)
	end := time.Now().Add(time.Minute * time.Duration(viper.GetInt("maxVideoProcessTimeInMin"))) // таймаут обработки видео
	for {
		time.Sleep(time.Second * 2)
		if time.Now().After(end) {
			fmt.Println("file processing timeout")
			break
		}

		item, err := h.services.File.GetByCode(userCode.String())
		if err != nil || item.Status == common.ERROR {
			fmt.Println("get file process status error:", err, item)
			break
		}

		if item.Status == common.FINISHED {
			err = ws.WriteJSON(item)
			fmt.Println("send data to websocket error, if exists:" + err.Error())
			break
		}
	}
}
