package test

import (
	"time"
	"tynmarket/coffeehub-go/model"
	"tynmarket/coffeehub-go/router"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Setup is
func Setup() *gorm.DB {
	db := model.TestDB()
	tx := model.StartTx(db)

	return tx
}

// TearDown is
func TearDown() {
	tx := model.TestDB()
	//tx.Commit()
	tx.Rollback()
}

// SetUpCoffees is
func SetUpCoffees() {
	db := model.TestDB()
	now := time.Now()

	site := model.Site{
		Name: "name",
		URL:  "url",
		Model: model.Model{
			CreatedAt: now, UpdatedAt: now,
		},
	}
	db.Create(&site)

	coffee := model.Coffee{
		SiteID:        site.ID,
		Path:          "/SHOP/CO-CY001.html",
		Country:       "コロンビア",
		AreaOrFactory: "フランコ・ロペス",
		Roast:         5,
		Taste:         "口に含んだ時のやわらかな食感とやさしいオレンジのような印象はこの地域の特徴です。心地よい軽めの濃縮感、飲みこんだ後には長い甘みの余韻が続きます。全てが高いレベルで調和しているコーヒーです。",
		Model: model.Model{
			CreatedAt: now, UpdatedAt: now,
		},
	}
	db.Create(&coffee)
}

// SetupRouter is
func SetupRouter() *gin.Engine {
	r := gin.Default()
	return router.Route(r)
}
