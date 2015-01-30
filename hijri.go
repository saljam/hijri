// Package hijri implement the Hijri tabular lunar calendar.
//
// It calculates the date using the following 30-year cyclic table:
//	354, 355, 354, 354, 355, 354, 355, 354, 354, 355,
//	354, 354, 355, 354, 354, 355, 354, 355, 354, 354,
//	355, 354, 354, 355, 354, 355, 354, 354, 355, 354,
package hijri

import "time"

// A Month specifies a Hijri month of the year (Muharram = 1, ...).
type Month int

const (
	Muharram Month = 1 + iota
	Safar
	RabiAlawwal
	RabiAlthani
	JumadaAlawwal
	JumadaAlthani
	Rajab
	Shaaban
	Ramadan
	Shawwal
	DhuAlQidah
	DhuAlHijjah
)

const (
	daysPer30Years   = 11*355 + 19*354
	daysPer2Months   = 29 + 30
	daysPerEvenMonth = 29
)

var (
	// The closest start of the 30-year cycle to unix 0.
	// 1 Muharram 1380, 26 June 1960
	zeroTime    = time.Unix(-300326400, 0)
	zeroYear    = 1380
	daysPerYear = []int{
		354, 355, 354, 354, 355, 354, 355, 354, 354, 355,
		354, 354, 355, 354, 354, 355, 354, 355, 354, 354,
		355, 354, 354, 355, 354, 355, 354, 354, 355, 354,
	}
)

// HijriDate returns the Hijri date for time t.
func HijriDate(t time.Time) (year int, month Month, day int) {
	// TODO don't use time.Duration because that's limited to 200something years.
	day = int(t.Sub(zeroTime).Hours()) / 24

	n := day / daysPer30Years
	year = zeroYear + 30*n
	day -= daysPer30Years * n

	for i := 0; day > daysPerYear[i]; i++ {
		year++
		day -= daysPerYear[i]
	}

	n = day / daysPer2Months
	month = Month(2 * n)
	day -= daysPer2Months * n
	if day > daysPerEvenMonth {
		month++
		day -= daysPerEvenMonth
	}

	month++
	day++
	return
}

// TODO Format dates, at least in English and Arabic.
