package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/subratkumarsahu/go-jwt/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/singup", controller.Signup())
	incomingRoutes.POST("users/login", controller.Login())
}
