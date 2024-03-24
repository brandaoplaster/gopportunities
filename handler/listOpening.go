package handler

import (
	"net/http"

	"github.com/brandaoplaster/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func ListOpeningsHandler(context *gin.Context) {
	openings := []schemas.Opening{}

	if error := db.Find(&openings).Error; error != nil {
		sendError(context, http.StatusInternalServerError, "internal server error")
		return
	}
	sendSuccess(context, "openings retrieved successfully", openings)
}
