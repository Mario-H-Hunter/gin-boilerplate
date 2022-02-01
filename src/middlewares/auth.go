package middlewares

import (
	"github.com/gin-gonic/gin"
	"incorta-cloud/upload-service/datasources"
	"incorta-cloud/upload-service/utils"
)

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		reqToken := c.Request.Header.Get("X-API-TOKEN")

		if len(reqToken) == 0 {
			c.AbortWithStatus(401)
		}

		//Temporarily set token for the backend to introspect
		c.Set(utils.UserToken, reqToken)

		var _, err = datasources.BackendClient{Context: c}.Get("/api/v1/users/introspect", nil)

		if err != nil {
			c.Set(utils.UserToken, nil)
			utils.UnauthorizedError(c, err)
			return
		}

		c.Next()
	}
}
