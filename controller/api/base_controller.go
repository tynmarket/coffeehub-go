package api

import (
	"fmt"
	"strconv"
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

func paginate(c *gin.Context, db *gorm.DB) *gorm.DB {
	page, _ := strconv.Atoi(c.Query("page"))
	return db.Offset(page * 10).Limit(11)
}
