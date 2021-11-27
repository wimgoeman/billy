package util

import "strconv"

type Percentage float32

func (percentage *Percentage) UnmarshalJSON(data []byte) error {
	rawData, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}
	perc64, err := strconv.ParseFloat(rawData, 32)
	*percentage = Percentage(perc64)
	return err
}

func (percentage Percentage) MarshalJSON() ([]byte, error) {
	return []byte(strconv.Quote(strconv.FormatFloat(float64(percentage), 'd', 5, 32))), nil
}

func (percentage Percentage) String() string {
	return strconv.FormatFloat(float64(percentage*100), 'f', -1, 32) + " %"
}
