package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type SendRequestBody struct {
	BotToken           string `json:"bot_token" binding:"required"`
	TextMessage        string `json:"text_message" binding:"required"`
	DestinationChannel string `json:"destination_channel" binding:"required"`
}

func handleSend() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody SendRequestBody
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"error":   err.Error(),
				"message": "You should provide bot_token, text_message and destination_channel",
			})
			return
		}

		bot, err := tgbotapi.NewBotAPI(requestBody.BotToken)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		message := tgbotapi.NewMessageToChannel(requestBody.DestinationChannel, requestBody.TextMessage)
		message.ParseMode = "markdown"

		if msgSent, err := bot.Send(message); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusOK, msgSent)
		}
	}
}

func handleIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
			"status":  "OK",
		},
		)
	}
}
