package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

/*
Нужно получить токен пользователя, провалидировать его и записать в контекст
*/

const (
	authorizationHandler = "Authorization"
	userCtx              = "userCode"
)

// userIdentity - проверка авторизации
func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHandler)
	if header == "" {
		NewErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if headerParts[0] != "Bearer" {
		NewErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if headerParts[1] == "" {
		NewErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	// чтобы каждый раз не ходить в БД за id пользователя по его коду -
	// в userCtx будем держать его uuid код
	userCode, err := h.services.User.ParseToken(headerParts[1])
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userCode)
}

// getUserCode - получение UserCode из контекста
func getUserCode(c *gin.Context) (string, error) {
	code, ok := c.Get(userCtx)
	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, "user code not found")
		return "", errors.New("user code not found")
	}

	userCodeString, ok := code.(uuid.UUID)
	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, "user code is of invalid type")
		return "", errors.New("user code is of invalid type")
	}

	return userCodeString.String(), nil
}
