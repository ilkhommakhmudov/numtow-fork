package kz

import (
	"github.com/ilkhommakhmudov/numtow-fork/curtow/cur"
)

const (
	minus       = "минус"
	zero        = "нөл"
	hundred     = "жүз"
	sep         = " "
	integerPart = "бүтін"
)

// nolint
var (
	megs  = [30]string{"", "мың", "миллион", "миллиард", "триллион", "квадриллион", "квинтиллион", "секстиллион", "сеплиллион", "октиллион", "нониллион", "дециллион", "ундециллион", "дуодециллион", "тредециллион", "кватуордециллион", "квиндециллион", "сексдециллион", "септендециллион", "октодециллион", "новемдециллион", "вигинтиллион", "унвигинтиллион", "дуовигинтиллион", "тревигинтиллион", "кватуорвигинтиллион", "квинвигинтиллион", "сексвигинтиллион", "септенвигинтиллион", "октовигинтиллион"}
	units = [10]string{"", "бір", "екі", "үш", "төрт", "бес", "алты", "жеті", "сегіз", "тоғыз"}
	tens  = [10]string{"", "он", "жиырма", "отыз", "қырық", "елу", "алпыс", "жетпіс", "сексен", "тоқсан"}

	fracPart = [90]string{
		"", "оннан", "жүзден",
		"мыңнан", "он мыңнан", "жүз мыңнан",
		"миллионнан", "он миллионнан", "жүз миллионнан",
		"миллиардтан", "он миллиардтан", "жүз миллиардтан",
		"триллионнан", "он триллионнан", "жүз триллионнан",
		"квадриллионнан", "он квадриллионнан", "жүз квадриллионнан",
		"квинтиллионнан", "он квинтиллионнан", "жүз квинтиллионнан",
		"секстиллионнан", "он секстиллионнан", "жүз секстиллионнан",
		"сеплиллионнан", "он сеплиллионнан", "жүз сеплиллионнан",
		"октиллионнан", "он октиллионнан", "жүз октиллионнан",
		"нониллионнан", "он нониллионнан", "жүз нониллионнан",
		"дециллионнан", "он дециллионнан", "жүз дециллионнан",
		"ундециллионнан", "он ундециллионнан", "жүз ундециллионнан",
		"дуодециллионнан", "он дуодециллионнан", "жүз дуодециллионнан",
		"тредециллионнан", "он тредециллионнан", "жүз тредециллионнан",
		"кватуордециллионнан", "он кватуордециллионнан", "жүз кватуордециллионнан",
		"квиндециллионнан", "он квиндециллионнан", "жүз квиндециллионнан",
		"сексдециллионнан", "он сексдециллионнан", "жүз сексдециллионнан",
		"септендециллионнан", "он септендециллионнан", "жүз септендециллионнан",
		"октодециллионнан", "он октодециллионнан", "жүз октодециллионнан",
		"новемдециллионнан", "он новемдециллионнан", "жүз новемдециллионнан",
		"вигинтиллионнан", "он вигинтиллионнан", "жүз вигинтиллионнан",
		"унвигинтиллионнан", "он унвигинтиллионнан", "жүз унвигинтиллионнан",
		"дуовигинтиллионнан", "он дуовигинтиллионнан", "жүз дуовигинтиллионнан",
		"тревигинтиллионнан", "он тревигинтиллионнан", "жүз тревигинтиллионнан",
		"кватуорвигинтиллионнан", "он кватуорвигинтиллионнан", "жүз кватуорвигинтиллионнан",
		"квинвигинтиллионнан", "он квинвигинтиллионнан", "жүз квинвигинтиллионнан",
		"сексвигинтиллионнан", "он сексвигинтиллионнан", "жүз сексвигинтиллионнан",
		"септенвигинтиллионнан", "он септенвигинтиллионнан", "жүз септенвигинтиллионнан",
		"октовигинтиллионнан", "он октовигинтиллионнан", "жүз октовигинтиллионнан"}

	curNamesKZ = map[cur.Currency]string{
		cur.KZT: "теңге",
		cur.RUB: "рубль",
		cur.USD: "АҚШ доллары",
		cur.EUR: "еуро",
	}

	curUnitNamesKZ = map[cur.Currency]string{
		cur.KZT: "тиын",
		cur.RUB: "тиын",
		cur.USD: "цент",
		cur.EUR: "евроцент",
	}
)
