package main

import (
	"fmt"
	"strconv"
)

type Address struct {
	AddressLine string
	City        string
	PostalCode  string
}

type Business struct {
	Name      string
	Address   Address
	VatNumber string
}

type Customer struct {
	Name      string
	Address   Address
	VatNumber string
}

type Price struct {
	value string
}

func (p Price) Float() float64 {
	val, err := strconv.ParseFloat(p.value, 64)
	if err != nil {
		panic(err)
	}
	return val
}

func NewPrice(value float64) Price {
	return Price{fmt.Sprintf("%.2f", value)}
}

type VatRegime struct {
	Name       string
	Percentage float32
}

type InvoiceLine struct {
	Description string
	Count       float32
	UnitPrice   Price
	Price       Price
	VatRegime   VatRegime
	Vat         Price
}

type Invoice struct {
	Customer   Customer
	TotalVat   Price
	TotalPrice Price
}
