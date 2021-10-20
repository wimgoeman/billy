package main

import "testing"

func TestPrice(t *testing.T) {
	price := Price{"125.25"}
	if price.Float() != 125.25 {
		t.Fatalf("Float value of Price was %f while 125.25 was expected", price.Float())
	}

	if NewPrice(125.25).value != "125.25" {
		t.Fatalf("String value of Price was %s while 125.25 was expected", NewPrice(125.25).value)
	}

	if NewPrice(125.256).value != "125.26" {
		t.Fatalf("String value of Price was %s while 125.26 was expected", NewPrice(125.25).value)
	}

	if NewPrice(125.2).value != "125.20" {
		t.Fatalf("String value of Price was %s while 125.20 was expected", NewPrice(125.25).value)
	}
}
