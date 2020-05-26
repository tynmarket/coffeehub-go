package router

import (
	"tynmarket/coffeehub-go/controller/api"

	"github.com/gin-gonic/gin"
)

// Route is
func Route(r *gin.Engine) *gin.Engine {
	apiV1 := r.Group("/api")
	apiV1.GET("/coffees", api.Coffees)
	apiV1.POST("/coffees", api.CoffeesCreate)
	apiV1.GET("/coffees/roast/:roast", api.CoffeesRoast)

	return r
}
