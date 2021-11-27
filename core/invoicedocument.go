package billy

type InvoiceDocumentCustomer struct {
	Name        string
	AddressLine string
	City        string
	PostalCode  string
	VatNumber   string
}

type InvoiceDocumentSelf struct {
	Name         string
	AddressLine  string
	City         string
	PostalCode   string
	VatNumber    string
	BankAccounts []string
}

type InvoiceDocumentLine struct {
	Description   string
	Count         float32
	UnitPrice     string
	PriceExclVat  string
	PriceInclVat  string
	VatPercentage string
	Vat           string
}

type InvoiceDocument struct {
	Customer     InvoiceDocumentCustomer
	Self         InvoiceDocumentSelf
	Number       int
	Date         string
	DueDate      string
	TotalExclVat string
	TotalInclVat string
	TotalVat     string
	Lines        []InvoiceDocumentLine
}
