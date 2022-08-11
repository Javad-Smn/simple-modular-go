package auth

import (
	"github.com/gin-gonic/gin"
)

func loginUser(c *gin.Context) {
	result := login(c)
	c.IndentedJSON(result.Status, result)
}
func signUp(c *gin.Context) {
	result := createUser(c)
	c.IndentedJSON(result.Status, result)
}
