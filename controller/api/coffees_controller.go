package api

import (
	"tynmarket/coffeehub-go/model"
	"tynmarket/coffeehub-go/serializer"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Coffees index action
func Coffees(c *gin.Context) {
	handle(c, func(db *gorm.DB) {
		coffees := []model.Coffee{}
		db = paginate(c, db)
		db.Preload("Site").Order("id desc").Find(&coffees)

		serialized := serializer.SerializeCoffees(coffees)

		c.JSON(200, serialized)

	})
}
