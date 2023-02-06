package zilla

type Thing struct {
	Name string `json:"name"`
}

type Image struct {
	ContentUrl string `json:"contentUrl"`
	Caption    string `json:"caption"`
}

type AggregateRating struct {
	WorstRating int    `json:"worstRating"`
	ReviewCount int    `json:"reviewCount"`
	RatingValue string `json:"ratingValue"`
	BestRating  int    `json:"bestRating"`
}

type Offer struct {
	Seller        Thing  `json:"seller"`
	PriceCurrency string `json:"priceCurrency"`
	Price         string `json:"price"`
	ItemCondition string `json:"itemCondition"`
	Availability  string `json:"availability"`
}

type Product struct {
	Sku             int             `json:"sku"`
	ProductID       int             `json:"productID"`
	Offers          Offer           `json:"offers"`
	Name            string          `json:"name"`
	Image           Image           `json:"image"`
	Description     string          `json:"description"`
	Color           string          `json:"color"`
	Category        string          `json:"category"`
	Brand           Thing           `json:"brand"`
	AggregateRating AggregateRating `json:"aggregateRating"`
}

type SkuDetails struct {
	AvailabilityLabel           string `json:"availability_label"`
	AvailabilityMessage         string `json:"availability_message"`
	Closeout                    bool   `json:"closeout"`
	Compound                    bool   `json:"compound"`
	Id                          int    `json:"id"`
	InStock                     bool   `json:"in_stock"`
	IsGuaranteedHolidayShipping bool   `json:"is_guaranteed_holiday_shipping"`
	IsPremiumEligible           bool   `json:"is_premium_eligible"`
	LoyaltyEarnings             string `json:"loyalty_earnings"`
	LoyaltyPrice                string `json:"loyalty_price"`
	LoyaltySavings              string `json:"loyalty_savings"`
	Msrp                        string `json:"msrp"`
	OptionsDescription          string `json:"options_description"`
	PercentOff                  int    `json:"percent_off"`
	Retail                      string `json:"retail"`
	SavingsAmount               string `json:"savings_amount"`
	ShowRetail                  bool   `json:"show_retail"`
}

type ParseResult struct {
	Products   []Product
	SkuDetails []SkuDetails
}
