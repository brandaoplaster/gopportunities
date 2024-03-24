package handler

import (
	"net/http"

	"github.com/brandaoplaster/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ShowOpeningHandler(context *gin.Context) {
	id := context.Query("id")

	if id == "" {
		sendError(context, http.StatusBadRequest, "id is required")
		return
	}

	opening := schemas.Opening{}

	if error := db.First(&opening, id).Error; error != nil {
		sendError(context, http.StatusNotFound, "opening not found")
		return
	}
	sendSuccess(context, "opening retrieved successfully", opening)
}
