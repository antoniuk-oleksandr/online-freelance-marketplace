package body

type Search struct {
	Query            *string  `query:"query"`
	Order            *int     `query:"order"`
	Sort             *int     `query:"sort"`
	Skill            []string `query:"skill"`
	Language         []string `query:"language"`
	Category         []string `query:"category"`
	PriceFrom        *float32 `query:"priceFrom"`
	PriceTo          *float32 `query:"priceTo"`
	RatingFrom       *float32 `query:"ratingFrom"`
	RatingTo         *float32 `query:"ratingTo"`
	DeliveryTimeFrom *float32 `query:"deliveryTimeFrom"`
	DeliveryTimeTo   *float32 `query:"deliveryTimeTo"`
	LevelFrom        *float32 `query:"levelFrom"`
	LevelTo          *float32 `query:"levelTo"`
	Cursor           *string  `query:"cursor"`
}
