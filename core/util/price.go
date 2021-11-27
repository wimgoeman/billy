package util

import (
	"fmt"
	"math"
	"strconv"
)

const (
	precision   = 2
	decimalSep  = ","
	thousandSep = " "
)

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

func (p Price) String() string {
	rounded := int64(math.Round(p.Float() * math.Pow10(precision)))
	bareformat := strconv.Itoa(int(rounded))

	// Add leading zeroes if needed
	for len(bareformat) < precision+1 {
		bareformat = "0" + bareformat
	}

	decimals := bareformat[len(bareformat)-precision:]
	thousands := bareformat[0 : len(bareformat)-precision]

	fullFormat := decimalSep + decimals

	for i := 0; i < len(thousands); i++ {
		character := thousands[len(thousands)-1-i]

		if i != 0 && (i%3) == 0 {
			fullFormat = " " + fullFormat
		}
		fullFormat = string(character) + fullFormat
	}

	return "€ " + fullFormat
}

func (p Price) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(p.value)), nil
}

func (p *Price) UnmarshalJSON(data []byte) error {
	rawData, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}
	*p = Price{rawData}
	return nil
}

func NewPrice(value float64) Price {
	return Price{fmt.Sprintf("%."+strconv.Itoa(precision)+"f", value)}
}
