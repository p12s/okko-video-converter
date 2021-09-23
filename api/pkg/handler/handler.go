package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/p12s/okko-video-converter/api/pkg/broker"
	"github.com/p12s/okko-video-converter/api/pkg/service"
)

type Handler struct {
	services *service.Service
	broker   *broker.Kafka
}

func NewHandler(services *service.Service, broker *broker.Kafka) *Handler {
	return &Handler{services: services, broker: broker}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	router.Use(CORSMiddleware())
	router.MaxMultipartMemory = 21 << 20 // 21 MiB (default is 32 MiB)

	router.Static("/files", "/files") // отдача статических файлов, в нашем случае - загр. файлов
	router.GET("/health", h.health)
	router.GET("/ws", h.websocket)
	router.POST("/api/registerUser", h.registerUser)

	api := router.Group("/api/v1", h.userIdentity)
	{
		api.GET("/files", h.files)
		api.POST("/upload", h.upload)
		api.GET("/removeAll", h.removeAll)
		api.GET("/download", h.download)
	}

	return router
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH,OPTIONS,GET,PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
