package controllers

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"

	"github.com/imthaghost/stockapi/stock"
)

// Company contains data about the company we are searching for
type Company struct {
	Ticker string `json:"ticker" form:"ticker" query:"ticker"`
}

// GetPrice handler method binds client JSON from body, form, or query string
// to the company strcutre and uses the ticker to crawl Yahoo Finance and return
// price formatted in JSON
func GetPrice(c echo.Context) (err error) {
	var bodyBytes []byte
	if c.Request().Body != nil {
		// Read the Body content
		bodyBytes, _ = ioutil.ReadAll(c.Request().Body)
	}
	// Restore the io.ReadCloser to its original state
	c.Request().Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	// Continue to use the Body, like Binding it to a struct:
	u := new(Company)
	// bind the model with the context body
	er := c.Bind(u)
	// panic!
	if er != nil {
		panic(err)
	}
	// crawl Yahoo Finance with such ticker
	r := stock.Price(u.Ticker)
	// Dictionary or Map for price and its data
	priceMap := map[string]string{"price": r}
	// JSON response
	return c.JSON(http.StatusOK, priceMap)
}
