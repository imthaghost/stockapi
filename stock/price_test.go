package stock

import "testing"

func BenchmarkFib10(b *testing.B) {
	Price("MSFT")
}
