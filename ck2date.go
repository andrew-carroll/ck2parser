package ck2save

import (
	"time"
)

type ck2Date struct {
	year  int
	month int
	day   int
}

func NewCK2Date(rawString string, quoted bool) ck2Date {
	switch quoted {
	case true:
		timeFormat := `"2006.1.2"`
		t, e := time.Parse(timeFormat, rawString)
		checkError(e)
		return ck2Date{t.Year(), int(t.Month()), t.Day()}
	default:
		timeFormat := "2006.1.2"
		t, e := time.Parse(timeFormat, rawString)
		checkError(e)
		return ck2Date{t.Year(), int(t.Month()), t.Day()}
	}
}
