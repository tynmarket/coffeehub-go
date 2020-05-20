package api

import (
	"github.com/gin-gonic/gin"
)

// Coffees index action
func Coffees(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "coffees",
	})
}
