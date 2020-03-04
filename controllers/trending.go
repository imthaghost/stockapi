package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/labstack/echo"

	"github.com/imthaghost/stockapi/stock"
)

// TrendingCompanies ...
type TrendingCompanies struct {
	name  string
	price string
}

const (
	apple      = "AAPL"
	google     = "GOOG"
	microsoft  = "MSFT"
	amazon     = "AMZN"
	facebook   = "FB"
	alibaba    = "BABA"
	intel      = "INTC"
	twitter    = "TWTR"
	salesforce = "CRM"
	intuit     = "INTU"
)

func prices() TrendingCompanies {
	//tickers := [10]string{apple, google, microsoft, amazon, facebook, alibaba, intel, twitter, salesforce, intuit}
	var example TrendingCompanies
	example = TrendingCompanies{"google", stock.Price("GOOG")}
	// var companyprices []TrendingCompanies
	// for _, value := range tickers {
	// 	x := TrendingCompanies{value, stock.Price(value)}
	// 	companyprices = append(companyprices, x)
	// }
	return example
}

// Trending handler
func Trending(c echo.Context) (err error) {
	x := prices()
	fmt.Println(x)
	something, _ := json.Marshal(x)
	fmt.Println(something)

	return c.JSON(200, x)
}
