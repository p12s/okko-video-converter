package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) registerUser(c *gin.Context) {
	// code - uuid-код созданного пользователя

	// возвр юид код
	userCode, err := h.services.CreateUser()
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// по коду генерир токен
	token, err := h.services.User.GenerateToken(userCode)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// отдаем
	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

}
