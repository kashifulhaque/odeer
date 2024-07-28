package server

import (
	"odeer/internal/config"
	"github.com/gin-gonic/gin"
)

func Start(cfg *config.Config) {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}