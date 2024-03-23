package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendError(context *gin.Context, statusCode int, message string) {
	context.Header("Content-Type", "application/json")
	context.JSON(statusCode, gin.H{"error": message})
}

func sendSuccess(context *gin.Context, message string, data interface{}) {
	context.Header("Content-Type", "application/json")
	context.JSON(http.StatusOK, gin.H{
		"message": message,
		"data":    data,
	})
}
