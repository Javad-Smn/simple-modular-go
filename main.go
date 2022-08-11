package main

import (
	"github.com/fatih/color"

	"github.com/Javad-Smn/simple-modular-go/src/libs"
	routers "github.com/Javad-Smn/simple-modular-go/src/modules"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	v1 := router.Group("v1")

	routers.V1(v1)
	color.HiYellow("Server started on port:" + libs.DotEnvVariable("PORT"))
	router.Run("localhost:" + libs.DotEnvVariable("PORT"))
}
