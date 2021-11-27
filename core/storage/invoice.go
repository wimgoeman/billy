package storage

import (
	"encoding/json"
	"time"

	"github.com/wimgoeman/billy/core/util"
)

type VatRegime struct {
	Code       string
	Percentage util.Percentage
}

type InvoiceLine struct {
	Description      string
	Count            float32
	UnitPriceExclVat util.Price
	VatRegime        VatRegime
}

type Invoice struct {
	CustomerId string
	Number     int
	DueDate    time.Time
	Date       time.Time
	Lines      []InvoiceLine
}

func ReadInvoice(id string) Invoice {
	data := util.OpenData(util.InvoiceDataType, id)
	defer data.Close()

	decoder := json.NewDecoder(data)
	invoice := Invoice{}
	err := decoder.Decode(&invoice)
	util.WarnOnErr(err)
	return invoice
}
