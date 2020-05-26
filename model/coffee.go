package model

import (
	"time"

	"github.com/go-playground/validator"
	"github.com/thoas/go-funk"
)

// Coffee is
type Coffee struct {
	Model
	SiteID        int64
	Site          Site
	Path          string `form:"path" binding:"required" validate:"required"`
	Country       string `form:"countory" binding:"required"`
	AreaOrFactory string `form:"area" binding:"required"`
	Roast         int    `form:"roast" binding:"required"`
	Taste         string `form:"taste" binding:"required" ja:"味わい"`
}

var coffeeRoastTextEn = []string{"unknown", "light", "cinnamon", "medium", "high", "city", "fullcity", "french", "italian"}
var coffeeRoastText = []string{"記載なし", "ライト", "シナモン", "ミディアム", "ハイ", "シティ", "フルシティ", "フレンチ", "イタリアン"}

// BeforeSave is
func (c *Coffee) BeforeSave() error {
	validate := validator.New()
	return validate.Struct(c)
}

// CoffeeRoastValue is
func CoffeeRoastValue(roast string) int {
	return funk.IndexOf(coffeeRoastTextEn, roast)
}

// ArrivalDate is
func (c *Coffee) ArrivalDate() string {
	return c.CreatedAt.Format("1月2日")
}

// ArrivalMonth is
func (c *Coffee) ArrivalMonth() int {
	month := c.CreatedAt.Month()
	return int(month)
}

// NewArrival is
func (c *Coffee) NewArrival() bool {
	t := time.Now()
	today := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return c.CreatedAt.After(today) || c.CreatedAt.Equal(today)
}

// RoastTextEn is
func (c *Coffee) RoastTextEn() string {
	return coffeeRoastTextEn[c.Roast]
}

// RoastText is
func (c *Coffee) RoastText() string {
	return coffeeRoastText[c.Roast]
}

// TasteText is
func (c *Coffee) TasteText() string {
	runeStr := []rune(c.Taste)

	if len(runeStr) > 120 {
		return string(runeStr[:120])
	}
	return c.Taste
}

// FullPath is
func (c *Coffee) FullPath() string {
	return c.Site.URL + c.Path
}
