package utils

import (
	"github.com/gin-gonic/gin"
)

func WriteResponse(c *gin.Context, status int, data interface{}) {
	c.JSON(status, data)
}

func ErrorResponse(c *gin.Context, status int, data interface{}, err error) {
	c.JSON(status, gin.H{
		"error": err.Error(),
		"data":  data,
	})
}
