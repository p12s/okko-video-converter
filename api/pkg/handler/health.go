package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Health
// @Tags health
// @Description Check service health
// @ID health
// @Produce  json
// @Success 200 {string} string OK
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /health [get]
func (h *Handler) health(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "OK🔥",
	})
}
