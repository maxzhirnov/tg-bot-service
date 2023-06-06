package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", handleIndex())
	r.POST("/send-text-to-channel", handleSend())

	r.Run(":8080")
}
