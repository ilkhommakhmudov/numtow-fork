package en

import (
	"errors"
	"testing"

	"github.com/ilkhommakhmudov/numtow-fork/curtow/cur"
)

//nolint:gochecknoglobals
var (
	testCaseString = []struct {
		GiveString    string
		GiveCurrency  cur.Currency
		GiveHideMU    bool
		GiveConvertMU bool
		WantResult    string
		WantError     error
	}{
		{GiveString: "1", GiveCurrency: cur.USD, GiveHideMU: true, WantResult: "one dollar"},
		{GiveString: "1.1", GiveCurrency: cur.USD, GiveConvertMU: true, WantResult: "one dollar and ten cents"},
		{GiveString: "1.0", WantResult: "one dollar and zero cents", GiveCurrency: cur.USD, GiveConvertMU: true},
	}
)

func TestCurrencyString(t *testing.T) {
	for _, v := range testCaseString {
		v := v

		t.Run(v.GiveString, func(tt *testing.T) {
			got, err := CurrencyString(v.GiveString, WithCur(v.GiveCurrency), WithCurIgnoreMU(v.GiveHideMU), WithCurConvMU(v.GiveConvertMU))
			if err != nil {
				if !errors.Is(v.WantError, err) {
					tt.Errorf("mismatch: want error: %s, wanr error: %s", err, v.WantError)
					return
				}
			}

			if got != v.WantResult {
				tt.Errorf("result mismatch: \nwant: %s\n got: %s", v.WantResult, got)
				return
			}
		})
	}
}
