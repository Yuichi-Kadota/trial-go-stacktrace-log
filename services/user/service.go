package user

import "github.com/gin-gonic/gin"

func get(c *gin.Context) error {
	r := NewUserRepo()
	err := r.GetUser()
	if err != nil {
		return err
	}
	return nil
}
