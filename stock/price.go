package stock

import (
	"github.com/antchfx/htmlquery"
)

// Price return the price of stock
func Price(ticker string) string {

	// price XPath
	pricePath := "//*[@id=\"quote-header-info\"]"
	// pattern := `USD+......`
	// re := regexp.MustCompile(pattern)

	// load the HTML
	doc, err := htmlquery.LoadURL(BaseURL + ticker)

	// uh oh :( freak out!!
	if err != nil {
		panic(err)
	}

	// HTML Node
	s2 := htmlquery.FindOne(doc, pricePath)

	// from the Node get inner next
	price := string(htmlquery.InnerText(s2))

	return price
}
