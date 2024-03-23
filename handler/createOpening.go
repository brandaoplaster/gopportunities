package handler

import (
	"net/http"

	"github.com/brandaoplaster/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

func CreateOpeningHandler(context *gin.Context) {
	request := CreateOpeningRequest{}

	context.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.ErrorFormat("validate error: %v", err.Error())
		sendError(context, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.ErrorFormat("CreateOpeningHandler: %v", err.Error())
		sendError(context, http.StatusInternalServerError, "internal server error")
		return
	}

	sendSuccess(context, "opening created successfully", opening)
}
