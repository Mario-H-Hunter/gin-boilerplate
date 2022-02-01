package repositories

import (
	"github.com/gin-gonic/gin"
	"incorta-cloud/upload-service/datasources"
	"incorta-cloud/upload-service/forms"
	"incorta-cloud/upload-service/models"
	"incorta-cloud/upload-service/utils"
)

type UserRepository struct {
}

func (r UserRepository) SignIn(c *gin.Context, signInForm forms.UserSignInRequest) (*models.UserSignInResponse, error) {

	var data, err = datasources.BackendClient{Context: c}.Post("api/v1/signin", nil, signInForm)

	if err != nil {
		return nil, err
	}

	var authBackendResponse models.SignInBackendResponse
	err = utils.ParseJson(data, &authBackendResponse)

	if err != nil || authBackendResponse.Status != 200 {
		return nil, err
	}

	return &models.UserSignInResponse{
		Token: authBackendResponse.Token,
		User:  *authBackendResponse.User,
	}, nil

}
