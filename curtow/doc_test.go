package curtow

import (
	"fmt"

	"github.com/ilkhommakhmudov/numtow-fork/curtow/cur"
	"github.com/ilkhommakhmudov/numtow-fork/lang"
	"github.com/ilkhommakhmudov/numtow-fork/lang/en"
)

func ExampleMustString_eur() {
	fmt.Println(MustString("12", lang.EN, en.WithCur(cur.EUR)))
	// Output: twelve euros and 00 cents
}

func ExampleMustString_usd() {
	fmt.Println(MustString("45.57", lang.EN, en.WithCur(cur.USD), en.WithCurConvMU(true)))
	// Output: forty-five dollars and fifty-seven cents
}
