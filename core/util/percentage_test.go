package util

import "testing"

func TestPercentageString(t *testing.T) {
	var testVal Percentage = 0.22
	assertEqual(t, testVal.String(), "22 %")
	testVal = 0.222
	assertEqual(t, testVal.String(), "22.2 %")
	testVal = 0
	assertEqual(t, testVal.String(), "0 %")
}
