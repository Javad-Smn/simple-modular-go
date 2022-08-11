package auth

import (
	"github.com/gin-gonic/gin"
)

func routes(rg *gin.RouterGroup) {
	authRoute := rg.Group("/auth")
	authRoute.POST("/login", loginUser)
	authRoute.POST("/sign-up", signUp)
}
func InitModule(rg *gin.RouterGroup) {
	routes(rg)
}
