package user

import (
	"encoding/json"
	"net/http"

	"github.com/Javad-Smn/simple-modular-go/src/libs"
	"github.com/Javad-Smn/simple-modular-go/src/structs"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func getUser(c *gin.Context) structs.Response {
	x, _ := c.Get("claims")
	var user structs.UserResponse
	json.Unmarshal([]byte(x.(jwt.MapClaims)["user"].(string)), &user)
	return libs.Response(http.StatusFound, "User found", user)
}
