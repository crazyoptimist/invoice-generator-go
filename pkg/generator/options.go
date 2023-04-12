package generator

// UnicodeTranslateFunc ...
type UnicodeTranslateFunc func(string) string

// Options for Document
type Options struct {
	AutoPrint bool `json:"auto_print,omitempty"`

	CurrencySymbol    string `default:"$ " json:"currency_symbol,omitempty"`
	CurrencyPrecision int    `default:"2" json:"currency_precision,omitempty"`
	CurrencyDecimal   string `default:"." json:"currency_decimal,omitempty"`
	CurrencyThousand  string `default:" " json:"currency_thousand,omitempty"`

	TextTypeInvoice      string `default:"INVOICE" json:"text_type_invoice,omitempty"`
	TextTypeQuotation    string `default:"QUOTATION" json:"text_type_quotation,omitempty"`
	TextTypeDeliveryNote string `default:"DELIVERY NOTE" json:"text_type_delivery_note,omitempty"`

	TextInvoiceDateTitle string `default:"Invoice Date" json:"text_invoice_date_title,omitempty"`
	TextDueDateTitle     string `default:"Due Date" json:"text_due_date_title,omitempty"`
	TextPaymentMethod    string `default:"Payment Method" json:"text_due_date_title,omitempty"`

	TextItemsNameTitle     string `default:"Description" json:"text_items_name_title,omitempty"`
	TextItemsUnitCostTitle string `default:"Unit price" json:"text_items_unit_cost_title,omitempty"`
	TextItemsQuantityTitle string `default:"Qty" json:"text_items_quantity_title,omitempty"`
	TextItemsTotalTitle    string `default:"Total" json:"text_items_total_ttc_title,omitempty"`

	TextTotalTotal string `default:"TOTAL" json:"text_total_total,omitempty"`

	BaseTextColor []int `default:"[35,35,35]" json:"base_text_color,omitempty"`
	GreyTextColor []int `default:"[82,82,82]" json:"grey_text_color,omitempty"`
	GreyBgColor   []int `default:"[232,232,232]" json:"grey_bg_color,omitempty"`
	DarkBgColor   []int `default:"[212,212,212]" json:"dark_bg_color,omitempty"`

	Font     string `default:"Helvetica"`
	BoldFont string `default:"Helvetica"`

	UnicodeTranslateFunc UnicodeTranslateFunc
}
