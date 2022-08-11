package auth

import (
	"net/http"

	libs "github.com/Javad-Smn/simple-modular-go/src/libs"
	userModule "github.com/Javad-Smn/simple-modular-go/src/modules/user"
	"github.com/Javad-Smn/simple-modular-go/src/structs"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func login(c *gin.Context) structs.Response {
	var body structs.User
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	password, err := userModule.FindUserPassword(body.Name)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(body.Password))
	if err != nil {
		c.AbortWithError(http.StatusForbidden, err)
	}

	user, _ := userModule.FindOne(body.Name)

	token, err := libs.GenerateJWT(user)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	return libs.Response(http.StatusAccepted, "You successfully logged in", token)
}
func createUser(c *gin.Context) structs.Response {
	var user structs.User
	if err := c.BindJSON(&user); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hashedPass)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	var insertedId = userModule.InsertOne(user)

	return libs.Response(http.StatusCreated, "User successfully created", insertedId)
}
