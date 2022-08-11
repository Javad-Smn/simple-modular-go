package libs

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Javad-Smn/simple-modular-go/src/structs"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var mySigningKey = []byte(DotEnvVariable("JWT_SECRET"))

func ValidateUserToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := jwt.Parse(c.Request.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("there was an error")
			}
			return mySigningKey, nil
		})
		if err != nil {
			c.Status(http.StatusUnauthorized)
			c.Writer.Write([]byte(err.Error()))
		}

		c.Set("claims", token.Claims)
		c.Next()
	}
}

func GenerateJWT(user structs.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["user"] = `{"id":"` + user.ID.Hex() + `","name":"` + user.Name + `","age":` + strconv.Itoa(int(user.Age)) + `}`
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
