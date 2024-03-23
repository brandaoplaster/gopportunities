package handler

import (
	"net/http"

	"github.com/brandaoplaster/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteOpeningHandler(context *gin.Context) {
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

	if error := db.Delete(&opening).Error; error != nil {
		sendError(context, http.StatusInternalServerError, "internal server error")
		return
	}

	sendSuccess(context, "opening deleted successfully", opening)
}
