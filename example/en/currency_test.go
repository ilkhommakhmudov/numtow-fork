package en

import (
	"fmt"
	"testing"

	"github.com/ilkhommakhmudov/numtow-fork/curtow"
	"github.com/ilkhommakhmudov/numtow-fork/curtow/cur"
	"github.com/ilkhommakhmudov/numtow-fork/lang"
	"github.com/ilkhommakhmudov/numtow-fork/lang/en"
)

func TestCurrencyEN(t *testing.T) {
	// convert currency to english words using curtow package
	fmt.Println(curtow.MustString("12", lang.EN, en.WithCur(cur.USD)))                           // twelve dollars and 00 cents
	fmt.Println(curtow.MustString("12", lang.EN, en.WithCur(cur.USD), en.WithCurIgnoreMU(true))) // twelve dollars
	fmt.Println(curtow.MustString("12", lang.EN, en.WithCur(cur.USD), en.WithCurConvMU(true)))   // twelve dollars and zero cents

	res, err := en.CurrencyString("53241", en.WithCur(cur.EUR))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(res) // fifty-three thousand, two hundred and forty-one euros and 00 cents
}
