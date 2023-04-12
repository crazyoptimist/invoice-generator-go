package generator

// SetType set type of document
func (d *Document) SetType(docType string) *Document {
	d.Type = docType
	return d
}

// SetHeader set header of document
func (d *Document) SetHeader(header *HeaderFooter) *Document {
	d.Header = header
	return d
}

// SetFooter set footer of document
func (d *Document) SetFooter(footer *HeaderFooter) *Document {
	d.Footer = footer
	return d
}

// SetNotes of document
func (d *Document) SetNotes(notes string) *Document {
	d.Notes = notes
	return d
}

// SetCompany of document
func (d *Document) SetCompany(company *Contact) *Document {
	d.Company = company
	return d
}

// SetCustomer of document
func (d *Document) SetCustomer(customer *Contact) *Document {
	d.Customer = customer
	return d
}

// AppendItem to document items
func (d *Document) AppendItem(item *Item) *Document {
	d.Items = append(d.Items, item)
	return d
}

// SetInvoiceDate of document
func (d *Document) SetInvoiceDate(date string) *Document {
	d.InvoiceDate = date
	return d
}

// SetDueDate of document
func (d *Document) SetDueDate(date string) *Document {
	d.DueDate = date
	return d
}

// Set PaymentMethod of document
func (d *Document) SetPaymentMethod(paymentMethod string) *Document {
	d.PaymentMethod = paymentMethod
	return d
}
