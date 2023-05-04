package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"

	"invoice-generator/pkg/generator"
)

const dateFormat = "01/02/2006"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(".env file is missing.")
	}

	invoiceDate := os.Getenv("INVOICE_DATE")
	if invoiceDate == "" {
		invoiceDate = time.Now().Format(dateFormat)
	}

	dueDate := os.Getenv("DUE_DATE")
	if dueDate == "" {
		dueDate = time.Now().AddDate(0, 0, 30).Format(dateFormat)
	}

	paymentMethod := os.Getenv("PAYMENT_METHOD")

	doc, _ := generator.New(generator.Invoice, &generator.Options{
		TextTypeInvoice: fmt.Sprintf("Invoice %s%s", os.Getenv("TITLE_PREFIX"), invoiceDate),
		AutoPrint:       true,
		CurrencySymbol:  "$",
	})

	doc.SetNotes("Thank you for your business!")

	doc.SetInvoiceDate(invoiceDate)
	doc.SetDueDate(dueDate)
	doc.SetPaymentMethod(paymentMethod)

	doc.SetCompany(&generator.Contact{
		Name: os.Getenv("COMPANY_NAME"),
		Address: &generator.Address{
			Address:    os.Getenv("COMPANY_ADDRESS"),
			Address2:   os.Getenv("COMPANY_ADDRESS_1"),
			PostalCode: os.Getenv("COMPANY_POSTAL_CODE"),
			City:       os.Getenv("COMPANY_CITY"),
			Country:    os.Getenv("COMPANY_COUNTRY"),
		},
	})

	doc.SetCustomer(&generator.Contact{
		Name: os.Getenv("CUSTOMER_NAME"),
		Address: &generator.Address{
			Address:    os.Getenv("CUSTOMER_ADDRESS"),
			Address2:   os.Getenv("CUSTOMER_ADDRESS_1"),
			PostalCode: os.Getenv("CUSTOMER_POSTAL_CODE"),
			City:       os.Getenv("CUSTOMER_CITY"),
			Country:    os.Getenv("CUSTOMER_COUNTRY"),
		},
	})

	doc.AppendItem(&generator.Item{
		Description: os.Getenv("SERVICE_DESCRIPTION"),
		UnitCost:    os.Getenv("UNIT_COST"),
		Quantity:    "1",
	})

	pdf, err := doc.Build()
	if err != nil {
		log.Fatal(err)
	}

	err = pdf.OutputFileAndClose(fmt.Sprintf("%s-invoice-%s.pdf", strings.Split(os.Getenv("COMPANY_NAME"), " ")[0], time.Now().Format("01-02-2006")))

	if err != nil {
		log.Fatal(err)
	}
}
