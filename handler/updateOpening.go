package handler

import (
	"net/http"

	"github.com/brandaoplaster/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func UpdateOpeningHandler(context *gin.Context) {
	request := UpdateOpeningRequest{}

	context.BindJSON(&request)

	if error := request.Validate(); error != nil {
		logger.ErrorFormat("error validating request: %s", error.Error())
		sendError(context, http.StatusBadRequest, error.Error())
		return
	}

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

	if request.Role != "" {
		opening.Role = request.Role
	}

	if request.Company != "" {
		opening.Company = request.Company
	}

	if request.Location != "" {
		opening.Location = request.Location
	}

	if request.Remote != nil {
		opening.Remote = *request.Remote
	}

	if request.Link != "" {
		opening.Link = request.Link
	}

	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	if error := db.Save(&opening).Error; error != nil {
		logger.ErrorFormat("error updating opening: %s", error.Error())
		sendError(context, http.StatusInternalServerError, "internal server error")
		return
	}

	sendSuccess(context, "opening updated successfully", opening)
}
