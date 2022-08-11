package routers

import (
	"github.com/Javad-Smn/simple-modular-go/src/modules/auth"
	"github.com/Javad-Smn/simple-modular-go/src/modules/user"
	"github.com/gin-gonic/gin"
)

func V1(r *gin.RouterGroup) {
	user.InitModule(r)
	auth.InitModule(r)
}
