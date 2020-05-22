package proto

import (
	"fmt"
	"tynmarket/coffeehub-go/model"

	"github.com/jinzhu/gorm"
)

func handle(handlerFun func(db *gorm.DB)) {
	db, err := model.Db()

	if err != nil {
		fmt.Println(err)
		return
	}

	defer db.Close()
	db.LogMode(true)

	handlerFun(db)
}
