package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/send-text-to-channel", handleSend())

	r.Run(":8080")
}
