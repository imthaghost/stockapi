package controllers

// Company contains data about the company we are searching for
type Company struct {
	Ticker string `json:"ticker" form:"ticker" query:"ticker"`
}
