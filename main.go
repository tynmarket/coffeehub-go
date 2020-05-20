package main

import (
	"tynmarket/coffeehub-go/controller/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	apiV1 := r.Group("/api")
	apiV1.GET("/coffees", api.Coffees)

	r.Run()
}
