package generator

import (
	"github.com/shopspring/decimal"
)

func (doc *Document) Total() decimal.Decimal {
	total := decimal.NewFromInt(0)

	for _, item := range doc.Items {
		total = total.Add(item.Total())
	}

	return total
}
