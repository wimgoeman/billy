package util

import (
	"encoding/json"
	"strings"
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

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

func TestPriceString(t *testing.T) {
	assertEqual(t, "1,65", NewPrice(1.651).String())
	assertEqual(t, "0,10", NewPrice(0.1).String())
	assertEqual(t, "9 999,00", NewPrice(9999).String())
	assertEqual(t, "9 999 999,00", NewPrice(9999999).String())
	assertEqual(t, "999,99", NewPrice(999.99).String())
	assertEqual(t, "0,00", NewPrice(0).String())
}

type TestStruct struct {
	Data Price
}

func TestJSON(t *testing.T) {
	builder := strings.Builder{}
	encoder := json.NewEncoder(&builder)
	err := encoder.Encode(TestStruct{Data: Price{"0.00"}})
	if err != nil {
		t.Fatalf("Failed to encode: %s", err)
	}
	expected := "{\"Data\":\"0.00\"}\n"
	assertEqual(t, builder.String(), expected)

	decoder := json.NewDecoder(strings.NewReader(expected))
	decoded := TestStruct{}
	err = decoder.Decode(&decoded)
	if err != nil {
		t.Fatalf("Failed to decode: %s", err)
	}
	assertEqual(t, decoded.Data.value, "0.00")
}
