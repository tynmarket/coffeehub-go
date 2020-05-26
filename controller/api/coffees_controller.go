package api

import (
	"fmt"
	"net/http"
	"time"
	"tynmarket/coffeehub-go/model"
	"tynmarket/coffeehub-go/serializer"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	"github.com/k0kubun/pp"
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

// CoffeesCreate create action
func CoffeesCreate(c *gin.Context) {
	handle(c, func(db *gorm.DB) {
		var coffee model.Coffee
		if err := c.ShouldBindWith(&coffee, binding.Form); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.(validator.ValidationErrors).Translate(trans)})
			return
		}

		coffee.SiteID = 1
		now := time.Now()
		coffee.CreatedAt = now
		coffee.UpdatedAt = now

		coffee.Path = ""
		pp.Printf("\ncoffee: %+v\n\n", &coffee)

		db = db.Create(&coffee)
		errors := db.GetErrors()

		if len(errors) > 0 {
			for _, err := range errors {
				fmt.Printf("\nerr: %s\n\n", err.Error())
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
		}

		c.JSON(200, "")
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
