package generator

import (
	"github.com/shopspring/decimal"
)

// Item represent a 'product' or a 'service'
type Item struct {
	Name        string `json:"name,omitempty" validate:"required"`
	Description string `json:"description,omitempty"`
	UnitCost    string `json:"unit_cost,omitempty"`
	Quantity    string `json:"quantity,omitempty"`

	_unitCost decimal.Decimal
	_quantity decimal.Decimal
}

// Prepare convert strings to decimal
func (i *Item) Prepare() error {
	// Unit cost
	unitCost, err := decimal.NewFromString(i.UnitCost)
	if err != nil {
		return err
	}
	i._unitCost = unitCost

	// Quantity
	quantity, err := decimal.NewFromString(i.Quantity)
	if err != nil {
		return err
	}
	i._quantity = quantity

	return nil
}

func (i *Item) Total() decimal.Decimal {
	quantity, _ := decimal.NewFromString(i.Quantity)
	price, _ := decimal.NewFromString(i.UnitCost)
	total := price.Mul(quantity)

	return total
}

// appendColTo document doc
func (i *Item) appendColTo(options *Options, doc *Document) {
	// Get base Y (top of line)
	baseY := doc.pdf.GetY()

	// Name
	doc.pdf.SetX(ItemColNameOffset)
	doc.pdf.MultiCell(
		ItemColUnitPriceOffset-ItemColNameOffset,
		3,
		doc.encodeString(i.Name),
		"",
		"",
		false,
	)

	// Description
	if len(i.Description) > 0 {
		doc.pdf.SetX(ItemColNameOffset)
		doc.pdf.SetY(doc.pdf.GetY() + 1)

		doc.pdf.SetFont(doc.Options.Font, "", SmallTextFontSize)
		doc.pdf.SetTextColor(
			doc.Options.GreyTextColor[0],
			doc.Options.GreyTextColor[1],
			doc.Options.GreyTextColor[2],
		)

		doc.pdf.MultiCell(
			ItemColUnitPriceOffset-ItemColNameOffset,
			3,
			doc.encodeString(i.Description),
			"",
			"",
			false,
		)

		// Reset font
		doc.pdf.SetFont(doc.Options.Font, "", BaseTextFontSize)
		doc.pdf.SetTextColor(
			doc.Options.BaseTextColor[0],
			doc.Options.BaseTextColor[1],
			doc.Options.BaseTextColor[2],
		)
	}

	// Compute line height
	colHeight := doc.pdf.GetY() - baseY

	// Unit price
	doc.pdf.SetY(baseY)
	doc.pdf.SetX(ItemColUnitPriceOffset)
	doc.pdf.CellFormat(
		ItemColQuantityOffset-ItemColUnitPriceOffset,
		colHeight,
		doc.encodeString(doc.ac.FormatMoneyDecimal(i._unitCost)),
		"0",
		0,
		"",
		false,
		0,
		"",
	)

	// Quantity
	doc.pdf.SetX(ItemColQuantityOffset)
	doc.pdf.CellFormat(
		ItemColTotalOffset-ItemColQuantityOffset,
		colHeight,
		doc.encodeString(i._quantity.String()),
		"0",
		0,
		"",
		false,
		0,
		"",
	)

	// TOTAL
	doc.pdf.SetX(ItemColTotalOffset)
	doc.pdf.CellFormat(
		190-ItemColTotalOffset,
		colHeight,
		doc.encodeString(doc.ac.FormatMoneyDecimal(i.Total())),
		"0",
		0,
		"",
		false,
		0,
		"",
	)

	// Set Y for next line
	doc.pdf.SetY(baseY + colHeight)
}
