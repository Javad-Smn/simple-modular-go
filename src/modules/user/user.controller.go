package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getMe(c *gin.Context) {
	c.JSON(http.StatusOK, getUser(c))
}
