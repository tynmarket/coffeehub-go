package proto

import (
	context "context"
	"fmt"
	"strconv"
	"tynmarket/coffeehub-go/model"
	"tynmarket/coffeehub-go/serializer"

	"github.com/jinzhu/gorm"
)

// CoffeeServerImpl is
type CoffeeServerImpl struct {
}

// GetCoffee is
func (*CoffeeServerImpl) GetCoffee(_ context.Context, req *GetCoffeeRequest) (*GetCoffeeResponse, error) {
	coffeeID, _ := strconv.Atoi(req.CoffeeId)
	var coffee serializer.Coffee

	handle(func(db *gorm.DB) {
		coffees := []model.Coffee{}
		db.Preload("Site").Where("id = ?", coffeeID).Order("id desc").Find(&coffees)

		serialized := serializer.SerializeCoffees(coffees)

		if len(serialized) != 0 {
			coffee = serialized[0]
			return
		}
		coffee = serializer.Coffee{}
	})

	name := fmt.Sprintf("%s %s", coffee.Country, coffee.Area)

	return &GetCoffeeResponse{
		Name: name,
	}, nil
}

// GetCoffees is
func (*CoffeeServerImpl) GetCoffees(_ context.Context, req *GetCoffeesRequest) (*GetCoffeesResponse, error) {
	var serialized []serializer.Coffee
	var resp []*Coffee

	handle(func(db *gorm.DB) {
		coffees := []model.Coffee{}
		db.Preload("Site").Order("id desc").Find(&coffees)

		serialized = serializer.SerializeCoffees(coffees)
	})

	if len(serialized) != 0 {
		for _, coffee := range serialized {
			resp = append(resp, &Coffee{
				Area:         coffee.Area,
				ArrivalDate:  coffee.ArrivalDate,
				ArrivalMonth: int32(coffee.ArrivalMonth),
				Country:      coffee.Country,
				New:          coffee.New,
				Roast:        coffee.Roast,
				RoastText:    coffee.RoastText,
				Shop:         coffee.Shop,
				Taste:        coffee.Taste,
				Url:          coffee.URL,
			})
		}
	}

	return &GetCoffeesResponse{
		Coffees: resp,
	}, nil
}
