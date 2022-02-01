package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"incorta-cloud/upload-service/config"
	"incorta-cloud/upload-service/middlewares"
	"os"
)

func Init() {
	configs := config.GetConfig()
	r := gin.Default()
	AddGuestRoutes(r)
	middlewares.AddMiddlewares(r)
	AddRoutes(r)
	err := r.Run(configs.GetString("server.port"))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
