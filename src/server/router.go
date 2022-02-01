package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	c "incorta-cloud/upload-service/controllers"
	_ "incorta-cloud/upload-service/docs"
)

func AddGuestRoutes(router *gin.Engine) {
	fmt.Println("Registering Guest Routes")
	healthC := new(c.HealthController)
	userC := new(c.UserController)

	//docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.GET("/health", healthC.Status)

	v1 := router.Group("/api/v1")
	{
		userGroup := v1.Group("/user")
		{
			userGroup.POST("/signin", userC.SignIn)
		}
	}
}

func AddRoutes(router *gin.Engine) {
	fmt.Println("Registering Auth Routes")

	// init controllers
	//userC := new(c.UserController)
	//
	//// use them for route actions
	//
	//v1 := router.Group("/api/v1")
	//{
	//
	//}
}
