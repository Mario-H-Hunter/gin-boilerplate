package controllers

import (
	"github.com/gin-gonic/gin"
	"incorta-cloud/upload-service/forms"
	"incorta-cloud/upload-service/models"
	"incorta-cloud/upload-service/repositories"
	"incorta-cloud/upload-service/utils"
)

type UserController struct{}

var userModel = new(models.User)

// SignIn
// @Summary Sign in to the upload service API
// @ID SignIn
// @Description Sign in to the service By delegating the request to the cloud BE
// @Accept  json
// @Produce  json
// @Tags	User
// @Param   payload body forms.UserSignInRequest true "User data"
// @Success 200 {object} models.UserSignInHttpResponse{} "ok"
// @Unauthorized 401 {object} models.HttpResponse{} "Unauthorized"
// @Router /user/signin [post]
func (u UserController) SignIn(c *gin.Context) {
	signIn := forms.UserSignInRequest{}
	err := utils.FillAndValidateForm(c, &signIn)

	if err != nil {
		return
	}

	resp, err := repositories.UserRepository{}.SignIn(c, signIn)

	if err != nil {
		utils.UnauthorizedError(c, err)
		return
	}

	utils.Success(c, resp)

}
