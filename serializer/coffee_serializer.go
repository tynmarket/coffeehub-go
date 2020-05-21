package serializer

import (
	"tynmarket/coffeehub-go/model"
)

// Coffee is
type Coffee struct {
	Area         string `json:"area"`
	ArrivalDate  string `json:"arrival_date"`
	ArrivalMonth int    `json:"arrival_month"`
	Country      string `json:"country"`
	New          bool   `json:"new"`
	Roast        string `json:"roast"`
	RoastText    string `json:"roast_text"`
	Shop         string `json:"shop"`
	Taste        string `json:"taste"`
	URL          string `json:"url"`
}

// SerializeCoffees is
func SerializeCoffees(coffees []model.Coffee) []Coffee {
	array := []Coffee{}

	for _, coffee := range coffees {
		serialized := serializeCoffee(coffee)
		array = append(array, serialized)
	}

	return array
}

func serializeCoffee(coffee model.Coffee) Coffee {
	serialized := Coffee{
		Area:    coffee.AreaOrFactory,
		Country: coffee.Country,
		Shop:    coffee.Site.Name,
	}

	serialized.ArrivalDate = coffee.ArrivalDate()
	serialized.ArrivalMonth = coffee.ArrivalMonth()
	serialized.New = coffee.NewArrival()
	serialized.Roast = coffee.RoastTextEn()
	serialized.RoastText = coffee.RoastText()
	serialized.Taste = coffee.TasteText()
	serialized.URL = coffee.FullPath()

	return serialized
}
