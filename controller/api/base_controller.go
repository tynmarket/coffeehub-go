package api

import (
	"fmt"
	"reflect"
	"strconv"
	"tynmarket/coffeehub-go/model"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/ja"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ja_translations "github.com/go-playground/validator/v10/translations/ja"
	"github.com/jinzhu/gorm"
)

var trans ut.Translator

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en.New()
		ja := ja.New()
		uni := ut.New(en, ja)
		// this is usually know or extracted from http 'Accept-Language' header
		// also see uni.FindTranslator(...)
		trans, _ = uni.GetTranslator("en")
		ja_translations.RegisterDefaultTranslations(v, trans)

		// フィールド
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			fieldName := fld.Tag.Get("ja")
			fmt.Printf("\nfieldName: %s\n\n", fieldName)
			if fieldName == "-" {
				return ""
			}
			return fieldName
		})
	}
}

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
