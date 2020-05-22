package api

import (
	"fmt"
	"tynmarket/coffeehub-go/model"
	"tynmarket/coffeehub-go/serializer"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Coffee bind params
type cffeeBind struct {
	Roast string `uri:"roast"`
}

// Coffees index action
func Coffees(c *gin.Context) {
	handle(c, func(db *gorm.DB) {
		coffees := []model.Coffee{}
		paginate(c, db).Preload("Site").Order("id desc").Find(&coffees)

		serialized := serializer.SerializeCoffees(coffees)

		c.JSON(200, serialized)
	})
}

// CoffeesRoast roast action
func CoffeesRoast(c *gin.Context) {
	handle(c, func(db *gorm.DB) {
		coffees := []model.Coffee{}
		coffeeRoast(c, paginate(c, db).Preload("Site").Order("id desc")).Find(&coffees)

		serialized := serializer.SerializeCoffees(coffees)

		c.JSON(200, serialized)
	})
}

func coffeeRoast(c *gin.Context, db *gorm.DB) *gorm.DB {
	var bind cffeeBind
	if err := c.ShouldBindUri(&bind); err != nil {
		fmt.Println(err)
		return db
	}
	roastValue := model.CoffeeRoastValue(bind.Roast)

	if roastValue != -1 {
		return db.Where("roast = ?", roastValue)
	}
	return db
}
