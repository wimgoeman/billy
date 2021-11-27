package storage

import (
	"encoding/json"

	"github.com/wimgoeman/billy/core/util"
)

type BusinessType string

const (
	BusinessTypeCustomer BusinessType = "Customer"
	BusinessTypeSupplier BusinessType = "Supplier"
	BusinessTypeSelf     BusinessType = "Self"
)

type Address struct {
	AddressLine string
	City        string
	PostalCode  string
}

type Business struct {
	Id           string
	Type         BusinessType
	Name         string
	Address      Address
	VatNumber    string
	BankAccounts []string
}

func ReadBusiness(id string) Business {
	data := util.OpenData(util.BusinessDataType, id)
	defer data.Close()

	decoder := json.NewDecoder(data)
	business := Business{}
	err := decoder.Decode(&business)
	util.WarnOnErr(err)
	return business
}
