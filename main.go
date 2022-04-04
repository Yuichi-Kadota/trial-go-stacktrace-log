package main

import (
	"trial-go-stacktrace/services/user"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/user", user.GetHandler())
	engine.Run(":8088")
}
