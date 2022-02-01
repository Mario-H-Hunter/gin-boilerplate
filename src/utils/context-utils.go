package utils

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"incorta-cloud/upload-service/models"
	"net/http"
	"strings"
)

func FillAndValidateForm(c *gin.Context, form interface{}) error {

	if validationErr := c.ShouldBindJSON(&form); validationErr != nil {
		ValidationError(c, validationErr)
		c.Abort()
		return validationErr
	}
	return nil
}

func ParseJson(data string, obj interface{}) error {
	r := strings.NewReader(data)
	var decoder = json.NewDecoder(r)

	return decoder.Decode(&obj)
}

func parseResponsePayload(appStatusCode models.AppStatusCode, message string, data interface{}) models.HttpResponse {
	return models.HttpResponse{
		Code:    appStatusCode,
		Message: message,
		Payload: data,
	}
}

func Success(c *gin.Context, obj interface{}) {
	c.JSON(http.StatusOK, parseResponsePayload(models.Success, "success", obj))
}

func Error(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusBadRequest, parseResponsePayload(models.GeneralError, "error", err))
}

func UnauthorizedError(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, parseResponsePayload(models.AuthorizationError, "Unauthorized", err))
}

func TokenExpiredError(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, parseResponsePayload(models.TokenExpiredError, "Token Expired", err))
}

func ValidationError(c *gin.Context, obj interface{}) {

	fmt.Println("validation Error", obj)
	c.AbortWithStatusJSON(http.StatusBadRequest, parseResponsePayload(models.ValidationError, "Validation Error!", obj))
}
