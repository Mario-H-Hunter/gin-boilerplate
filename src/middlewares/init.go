package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var middlewares = []gin.HandlerFunc{
	authMiddleware(),
	corsMiddleware(),
}

func AddMiddlewares(r *gin.Engine) {
	fmt.Println("Registering Middlewares")

	r.Use(middlewares...)
}
