package generator

import (
	"fmt"

	"github.com/go-pdf/fpdf"
)

// Build pdf document from data provided
func (doc *Document) Build() (*fpdf.Fpdf, error) {
	// Validate document data
	if err := doc.Validate(); err != nil {
		return nil, err
	}

	// Build base doc
	doc.pdf.SetMargins(BaseMargin, BaseMarginTop, BaseMargin)
	doc.pdf.SetXY(10, 10)
	doc.pdf.SetTextColor(
		doc.Options.BaseTextColor[0],
		doc.Options.BaseTextColor[1],
		doc.Options.BaseTextColor[2],
	)

	// Set header
	if doc.Header != nil {
		if err := doc.Header.applyHeader(doc); err != nil {
			return nil, err
		}
	}

	// Set footer
	if doc.Footer != nil {
		if err := doc.Footer.applyFooter(doc); err != nil {
			return nil, err
		}
	}

	// Add first page
	doc.pdf.AddPage()

	// Load font
	doc.pdf.SetFont(doc.Options.Font, "", 12)

	// Appenf document title
	doc.appendTitle()

	// Append company contact to doc
	companyBottom := doc.Company.appendCompanyContactToDoc(doc)

	// Append customer contact to doc
	customerBottom := doc.Customer.appendCustomerContactToDoc(doc)

	if customerBottom > companyBottom {
		doc.pdf.SetXY(10, customerBottom)
	} else {
		doc.pdf.SetXY(10, companyBottom)
	}

	// Append items
	doc.appendItems()

	// Check page height (total bloc height = 30, 45 when doc discount)
	offset := doc.pdf.GetY() + 30
	if offset > MaxPageHeight {
		doc.pdf.AddPage()
	}

	// Append notes
	doc.appendNotes()

	// Append total
	doc.appendTotal()

	// Append payment term
	doc.appendPaymentTerm()

	// Append js to autoprint if AutoPrint == true
	if doc.Options.AutoPrint {
		doc.pdf.SetJavascript("print(true);")
	}

	return doc.pdf, nil
}

// appendTitle to document
func (doc *Document) appendTitle() {
	title := doc.typeAsString()

	// Set x y
	x, _, _, _ := doc.pdf.GetMargins()
	doc.pdf.SetXY(x, BaseMarginTop)

	// Draw rect
	doc.pdf.SetFillColor(doc.Options.DarkBgColor[0], doc.Options.DarkBgColor[1], doc.Options.DarkBgColor[2])
	doc.pdf.Rect(x, BaseMarginTop, 190, 10, "F")

	// Draw text
	doc.pdf.SetFont(doc.Options.Font, "", 14)
	doc.pdf.CellFormat(190, 10, doc.encodeString(title), "0", 0, "C", false, 0, "")
}

// drawTableTitles in document
func (doc *Document) drawTableTitles() {
	// Draw table titles
	doc.pdf.SetX(10)
	doc.pdf.SetY(doc.pdf.GetY() + 6)
	doc.pdf.SetFont(doc.Options.BoldFont, "B", 8)

	// Draw rec
	doc.pdf.SetFillColor(doc.Options.GreyBgColor[0], doc.Options.GreyBgColor[1], doc.Options.GreyBgColor[2])
	doc.pdf.Rect(10, doc.pdf.GetY(), 190, 6, "F")

	// Name
	doc.pdf.SetX(ItemColNameOffset)
	doc.pdf.CellFormat(
		ItemColUnitPriceOffset-ItemColNameOffset,
		6,
		doc.encodeString(doc.Options.TextItemsNameTitle),
		"0",
		0,
		"",
		false,
		0,
		"",
	)

	// Unit price
	doc.pdf.SetX(ItemColUnitPriceOffset)
	doc.pdf.CellFormat(
		ItemColQuantityOffset-ItemColUnitPriceOffset,
		6,
		doc.encodeString(doc.Options.TextItemsUnitCostTitle),
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
		6,
		doc.encodeString(doc.Options.TextItemsQuantityTitle),
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
		6,
		doc.encodeString(doc.Options.TextItemsTotalTitle),
		"0",
		0,
		"",
		false,
		0,
		"",
	)
}

// appendItems to document
func (doc *Document) appendItems() {
	doc.drawTableTitles()

	doc.pdf.SetX(10)
	doc.pdf.SetY(doc.pdf.GetY() + 8)
	doc.pdf.SetFont(doc.Options.Font, "", 8)

	for i := 0; i < len(doc.Items); i++ {
		item := doc.Items[i]

		// Append to pdf
		item.appendColTo(doc.Options, doc)

		if doc.pdf.GetY() > MaxPageHeight {
			// Add page
			doc.pdf.AddPage()
			doc.drawTableTitles()
			doc.pdf.SetFont(doc.Options.Font, "", 8)
		}

		doc.pdf.SetX(10)
		doc.pdf.SetY(doc.pdf.GetY() + 6)
	}
}

// appendNotes to document
func (doc *Document) appendNotes() {
	if len(doc.Notes) == 0 {
		return
	}

	currentY := doc.pdf.GetY()

	doc.pdf.SetFont(doc.Options.Font, "", 9)
	doc.pdf.SetX(BaseMargin)
	doc.pdf.SetRightMargin(100)
	doc.pdf.SetY(currentY + 10)

	_, lineHt := doc.pdf.GetFontSize()
	html := doc.pdf.HTMLBasicNew()
	html.Write(lineHt, doc.encodeString(doc.Notes))

	doc.pdf.SetRightMargin(BaseMargin)
	doc.pdf.SetY(currentY)
}

// appendTotal to document
func (doc *Document) appendTotal() {
	doc.pdf.SetY(doc.pdf.GetY() + 10)
	doc.pdf.SetFont(doc.Options.Font, "", LargeTextFontSize)
	doc.pdf.SetTextColor(
		doc.Options.BaseTextColor[0],
		doc.Options.BaseTextColor[1],
		doc.Options.BaseTextColor[2],
	)

	// Draw TOTAL HT title
	doc.pdf.SetX(120)
	doc.pdf.SetFillColor(doc.Options.DarkBgColor[0], doc.Options.DarkBgColor[1], doc.Options.DarkBgColor[2])
	doc.pdf.Rect(120, doc.pdf.GetY(), 40, 10, "F")
	doc.pdf.CellFormat(38, 10, doc.encodeString(doc.Options.TextTotalTotal), "0", 0, "R", false, 0, "")

	// Draw TOTAL HT amount
	doc.pdf.SetX(162)
	doc.pdf.SetFillColor(doc.Options.GreyBgColor[0], doc.Options.GreyBgColor[1], doc.Options.GreyBgColor[2])
	doc.pdf.Rect(160, doc.pdf.GetY(), 40, 10, "F")
	doc.pdf.CellFormat(
		40,
		10,
		doc.encodeString(doc.ac.FormatMoneyDecimal(doc.Total())),
		"0",
		0,
		"L",
		false,
		0,
		"",
	)

	// doc.pdf.SetY(doc.pdf.GetY() + 10)
}

// appendPaymentTerm to document
func (doc *Document) appendPaymentTerm() {
	if len(doc.InvoiceDate) > 0 {
		invoiceDateString := fmt.Sprintf(
			"%s: %s",
			doc.encodeString(doc.Options.TextInvoiceDateTitle),
			doc.encodeString(doc.InvoiceDate),
		)
		doc.pdf.SetY(doc.pdf.GetY() + 15)

		doc.pdf.SetX(120)
		doc.pdf.SetFont(doc.Options.BoldFont, "B", 10)
		doc.pdf.CellFormat(80, 4, doc.encodeString(invoiceDateString), "0", 0, "R", false, 0, "")
	}
	if len(doc.DueDate) > 0 {
		dueDateString := fmt.Sprintf(
			"%s: %s",
			doc.encodeString(doc.Options.TextDueDateTitle),
			doc.encodeString(doc.DueDate),
		)
		doc.pdf.SetY(doc.pdf.GetY() + 5)

		doc.pdf.SetX(120)
		doc.pdf.SetFont(doc.Options.BoldFont, "B", 10)
		doc.pdf.CellFormat(80, 4, doc.encodeString(dueDateString), "0", 0, "R", false, 0, "")
	}
	if len(doc.PaymentMethod) > 0 {
		paymentMethodString := fmt.Sprintf(
			"%s: %s",
			doc.encodeString(doc.Options.TextPaymentMethod),
			doc.encodeString(doc.PaymentMethod),
		)
		doc.pdf.SetY(doc.pdf.GetY() + 5)

		doc.pdf.SetX(120)
		doc.pdf.SetFont(doc.Options.BoldFont, "B", 10)
		doc.pdf.CellFormat(80, 4, doc.encodeString(paymentMethodString), "0", 0, "R", false, 0, "")
	}
}
