package user

import (
	"fmt"
	"net/http"
	"trial-go-stacktrace/internal/logger"

	"github.com/gin-gonic/gin"
)

func GetHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger.Info("auth/get")
		err := get(c)
		if err != nil {
			//ログ出力
			logger.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintln(err),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"user": "userInfo",
		})
	}
}
