package model

type Company struct {
	GUID        string        `gorm:"primaryKey" json:"guid"`
	CompanyName string        `json:"company_name"`
	ShortName   string        `json:"short_name"`
	SymbolCode  string        `json:"symbol_code"`
	Address     string        `json:"address"`
	Info        string        `json:"info"`
	MarketGUID  string        `json:"market_guid"`
	SectorGUID  string        `json:"sector_guid"`
	TodayPoints []TodayPoints `gorm:"foreignKey:CompanyGUID" json:"today_points"`
	IsFavorite  *bool         `json:"is_favorite"`
}

type TodayPoints struct {
	ID                        uint    `gorm:"primaryKey"`
	CompanyGUID               string  `gorm:"index;not null" json:"company_guid"`
	TradeDate                 string  `json:"trade_date"`
	Open                      float64 `json:"open"`
	Close                     float64 `json:"close"`
	High                      float64 `json:"high"`
	Low                       float64 `json:"low"`
	PreviousClose             float64 `json:"previous_close"`
	ChangeValue               float64 `json:"change_value"`
	ChangeRatio               float64 `json:"change_ratio"`
	TradeCount                float64 `json:"trade_count"`
	TradeValue                float64 `json:"trade_value"`
	TradeVolume               float64 `json:"trade_volume"`
	FiftyTwoWeekHigh          float64 `json:"fifty_two_week_high"`
	FiftyTwoWeekLow           float64 `json:"fifty_two_week_low"`
	PreviousClose7DaysBack    float64 `json:"previous_close_7_days_back"`
	PreviousClose30DaysBack   float64 `json:"previous_close_30_days_back"`
	PreviousClose365DaysBack  float64 `json:"previous_close_365_days_back"`
	ReturnValueLastWeek       float64 `json:"return_value_last_week"`
	ReturnValueLastMonth      float64 `json:"return_value_last_month"`
	ReturnValueLastYear       float64 `json:"return_value_last_year"`
	ReturnLastWeek            float64 `json:"return_last_week"`
	ReturnLastMonth           float64 `json:"return_last_month"`
	ReturnLastYear            float64 `json:"return_last_year"`
	DailyPriceToEarnings      float64 `json:"daily_price_to_earnings"`
	DailyPriceToBookValue     float64 `json:"daily_price_to_book_value"`
	DailyMarketCapitalization float64 `json:"daily_market_capitalization"`
	BasicEarningsPerShareTTM  float64 `json:"basic_earnings_per_share_ttm"`
	BookValuePerShareTTM      float64 `json:"book_value_per_share_ttm"`
}
