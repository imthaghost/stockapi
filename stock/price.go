package stock

import (
	"log"

	"github.com/antchfx/htmlquery"
)

const (
	// yahoo base URL
	baseURL = "https://finance.yahoo.com/quote/"
	// price XPath
	pricePath = "//*[@id=\"quote-header-info\"]"
)

// equivalent to panic
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Price return the price of stock
func Price(ticker string) string {
	// load the HTML
	doc, err := htmlquery.LoadURL(baseURL + ticker)
	// uh oh :( freak out!!
	check(err)
	// HTML Node
	s2 := htmlquery.FindOne(doc, pricePath)
	// from the Node get inner next
	price := string(htmlquery.InnerText(s2))
	return price
}
