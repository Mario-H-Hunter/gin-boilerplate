package utils

import (
	"github.com/gin-gonic/gin"
)
import "github.com/golang-jwt/jwt"

func GetClaims(c *gin.Context) (jwt.MapClaims, error) {
	var token = c.GetString(UserToken)

	jwtToken, err := jwt.Parse(token, nil)

	if err != nil && err.(*jwt.ValidationError).Errors != 2 {
		return nil, err
	}

	claims := jwtToken.Claims.(jwt.MapClaims)
	return claims, nil
}

func GetClaimedInstances(c *gin.Context) (interface{}, error) {
	claims, err := GetClaims(c)
	if err != nil {
		return nil, err
	}
	return claims["instances"], nil

}

func GetUserId(c *gin.Context) (string, error) {
	claims, err := GetClaims(c)
	if err != nil {
		return "", err
	}

	return claims["userId"].(string), nil

}
