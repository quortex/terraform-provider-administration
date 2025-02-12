package client

type LimitsItem struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type PrincingItem struct {
	SubscribeForYear     int     `json:"subscribe_for_year"`
	MonthlyPrice         float64 `json:"monthly_price"`
	MonthlyPriceCurrency string  `json:"monthly_price_currency"`
}

type Plan struct {
	ID       int            `json:"id,omitempty"`
	Name     string         `json:"name"`
	Features []string       `json:"features"`
	Limits   []LimitsItem   `json:"limits"`
	Pricing  []PrincingItem `json:"pricing"`
}
