package api

import (
	"fmt"
	"tynmarket/coffeehub-go/model"
	"tynmarket/coffeehub-go/serializer"

	"github.com/gin-gonic/gin"
)

// Coffees index action
func Coffees(c *gin.Context) {
	db, err := model.Db()

	if err != nil {
		fmt.Println(err)
		//os.Exit(1)
		c.JSON(200, gin.H{
			"message": "error",
		})
		return
	}

	defer db.Close()
	db.LogMode(true)

	coffees := []model.Coffee{}
	db.Preload("Site").Find(&coffees)

	serialized := serializer.SerializeCoffees(coffees)

	c.JSON(200, serialized)
}
