package api

import (
	"fmt"
	"tynmarket/coffeehub-go/model"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func handle(c *gin.Context, handlerFun func(db *gorm.DB)) {
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

	handlerFun(db)
}
