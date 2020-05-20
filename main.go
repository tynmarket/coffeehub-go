package main

import (
	"tynmarket/coffeehub-go/controllers/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	api_v1 := r.Group("/api")
	api_v1.GET("/coffees", api.Coffees)

	r.Run()
}
