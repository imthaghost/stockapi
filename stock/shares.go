package stock

import "github.com/antchfx/htmlquery"

// Shares ...
func Shares(ticker string) string {
	doc, err := htmlquery.LoadURL("https://finance.yahoo.com/quote/" + ticker)
	if err != nil {
		panic(err)
	}
	path := "//*[@id=\"quote-header-info\"]"
	//alternative method,but simple more.
	s2 := htmlquery.FindOne(doc, path)
	price := string(htmlquery.InnerText(s2))
	return price
}