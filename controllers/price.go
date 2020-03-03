package controllers

import (
	"bytes"
	"io/ioutil"

	"github.com/labstack/echo"

	"github.com/imthaghost/stockapi/stock"
)

//Price is data about the site we will scrape
type Price struct {
	Ticker string `json:"ticker" form:"ticker" query:"ticker"`
}

//GetPrice ...
func GetPrice(c echo.Context) (err error) {
	// Read the Body content
	var bodyBytes []byte
	if c.Request().Body != nil {
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	}
	// Restore the io.ReadCloser to its original state
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	// Continue to use the Body, like Binding it to a struct:
	u := new(Price)
	// bind the model with the context body
	er := c.Bind(u)
	// panic!
	if er != nil {
		panic(err)
	}
	// crawl with the passed in data
	r := stock.Price(u.Ticker)
	// return the links
	c.Logger().Print(r)
	return c.JSON(200, r)
}
