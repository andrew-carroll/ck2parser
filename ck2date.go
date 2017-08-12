package ck2save

import (
	"time"
)

type CK2Date struct {
	year  int
	month time.Month
	day   int
}

func NewCK2Date(rawString string, quoted bool) CK2Date {
	switch quoted {
	case true:
		timeFormat := `"2006.1.2"`
		t, e := time.Parse(timeFormat, rawString)
		checkError(e)
		return CK2Date{t.Year(), t.Month(), t.Day()}
	default:
		timeFormat := "2006.1.2"
		t, e := time.Parse(timeFormat, rawString)
		checkError(e)
		return CK2Date{t.Year(), t.Month(), t.Day()}
	}
}
