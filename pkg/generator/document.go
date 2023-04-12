package generator

import (
	"github.com/go-pdf/fpdf"
	"github.com/leekchan/accounting"
)

// Document define base document
type Document struct {
	pdf *fpdf.Fpdf
	ac  accounting.Accounting

	Options       *Options      `json:"options,omitempty"`
	Header        *HeaderFooter `json:"header,omitempty"`
	Footer        *HeaderFooter `json:"footer,omitempty"`
	Type          string        `json:"type,omitempty" validate:"required,oneof=INVOICE DELIVERY_NOTE QUOTATION"`
	Version       string        `json:"version,omitempty" validate:"max=32"`
	ClientRef     string        `json:"client_ref,omitempty" validate:"max=64"`
	Description   string        `json:"description,omitempty" validate:"max=1024"`
	Notes         string        `json:"notes,omitempty"`
	Company       *Contact      `json:"company,omitempty" validate:"required"`
	Customer      *Contact      `json:"customer,omitempty" validate:"required"`
	Items         []*Item       `json:"items,omitempty"`
	ValidityDate  string        `json:"validity_date,omitempty"`
	InvoiceDate   string        `json:"invoice_date,omitempty"`
	DueDate       string        `json:"due_date,omitempty"`
	PaymentMethod string        `json:"payment_method,omitempty"`
}

// Pdf returns the underlying *fpdf.Fpdf used to build document
func (doc *Document) Pdf() *fpdf.Fpdf {
	return doc.pdf
}

// SetUnicodeTranslator to use
// See https://pkg.go.dev/github.com/go-pdf/fpdf#UnicodeTranslator
func (doc *Document) SetUnicodeTranslator(fn UnicodeTranslateFunc) {
	doc.Options.UnicodeTranslateFunc = fn
}

// encodeString encodes the string using doc.Options.UnicodeTranslateFunc
func (doc *Document) encodeString(str string) string {
	return doc.Options.UnicodeTranslateFunc(str)
}

// typeAsString return the document type as string
func (d *Document) typeAsString() string {
	if d.Type == Invoice {
		return d.Options.TextTypeInvoice
	}

	if d.Type == Quotation {
		return d.Options.TextTypeQuotation
	}

	return d.Options.TextTypeDeliveryNote
}
