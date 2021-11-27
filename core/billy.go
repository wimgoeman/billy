package billy

import (
	"errors"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/wimgoeman/billy/core/storage"
	"github.com/wimgoeman/billy/core/util"
)

func SendInvoice(templateName string, invoiceMatcher string) (string, error) {
	invoice := loadInvoice(invoiceMatcher)

	templateInfo, err := GetTemplate(templateName)
	if err != nil {
		return "", err
	}

	targetRoot, err := os.MkdirTemp("", "invoice")
	if err != nil {
		return "", err
	}
	err = filepath.WalkDir(templateInfo.Path, func(walkpath string, d fs.DirEntry, err error) error {
		if !strings.HasPrefix(walkpath, templateInfo.Path) {
			panic(errors.New("Walked path does not have the paths' prefix: " + walkpath))
		}

		relative := walkpath[len(templateInfo.Path):]
		if relative == "" {
			return nil
		}
		fullTarget := path.Join(targetRoot, relative)

		if !d.IsDir() {
			templateSource, err := os.Open(walkpath)
			if err != nil {
				return err
			}
			defer templateSource.Close()

			bytes, err := ioutil.ReadAll(templateSource)
			if err != nil {
				return err
			}

			template, err := template.New("invoice").Parse(string(bytes))
			if err != nil {
				return err
			}

			targetFile, err := os.Create(fullTarget)
			if err != nil {
				return err
			}

			defer targetFile.Close()
			err = template.Execute(targetFile, invoice)
			if err != nil {
				return err
			}
		} else {
			err := os.Mkdir(fullTarget, os.ModePerm)
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	return path.Join(targetRoot, "index.html"), nil
}

func loadInvoice(id string) InvoiceDocument {
	// TODO: Nice mechanism to "find" an invoice based on input
	invoiceData := storage.ReadInvoice(id)
	customer := storage.ReadBusiness(invoiceData.CustomerId)

	// TODO: Properly find out "self" business ID
	self := storage.ReadBusiness("1")

	document := InvoiceDocument{
		Customer: InvoiceDocumentCustomer{
			Name:        customer.Name,
			VatNumber:   customer.VatNumber,
			AddressLine: customer.Address.AddressLine,
			City:        customer.Address.City,
			PostalCode:  customer.Address.PostalCode,
		},
		Self: InvoiceDocumentSelf{
			Name:         self.Name,
			City:         self.Address.City,
			AddressLine:  self.Address.AddressLine,
			PostalCode:   self.Address.PostalCode,
			VatNumber:    self.VatNumber,
			BankAccounts: self.BankAccounts,
			ContactEmail: self.ContactEmail,
			ContactName:  self.ContactName,
			ContactPhone: self.ContactPhone,
		},
		Lines:   make([]InvoiceDocumentLine, len(invoiceData.Lines)),
		Number:  invoiceData.Number,
		Date:    invoiceData.Date.Format("02 Jan 2006"),
		DueDate: invoiceData.DueDate.Format("02 Jan 2006"),
	}

	var totalVat util.Price = util.NewPrice(0)
	var totalInclVat util.Price = util.NewPrice(0)
	var totalExclVat util.Price = util.NewPrice(0)

	for i, lineData := range invoiceData.Lines {
		lineExclVat := util.NewPrice(lineData.UnitPriceExclVat.Float() * float64(lineData.Count))
		lineVat := util.NewPrice(lineExclVat.Float() * float64(lineData.VatRegime.Percentage))
		lineInclVat := util.NewPrice(lineExclVat.Float() + lineVat.Float())

		docLine := &document.Lines[i]
		docLine.Count = lineData.Count
		docLine.Description = lineData.Description
		docLine.PriceExclVat = lineExclVat.String()
		docLine.PriceInclVat = lineInclVat.String()
		docLine.UnitPrice = lineData.UnitPriceExclVat.String()
		docLine.Vat = lineVat.String()
		docLine.VatPercentage = lineData.VatRegime.Percentage.String()

		totalVat = util.NewPrice(totalVat.Float() + lineVat.Float())
		totalInclVat = util.NewPrice(totalInclVat.Float() + lineInclVat.Float())
		totalExclVat = util.NewPrice(totalExclVat.Float() + lineExclVat.Float())
	}

	document.TotalExclVat = totalExclVat.String()
	document.TotalInclVat = totalInclVat.String()
	document.TotalVat = totalVat.String()

	return document
}
