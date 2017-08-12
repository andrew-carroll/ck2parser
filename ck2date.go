package ck2save

import (
	"time"
)

type CK2Date struct {
	year  int
	month time.Month
	day   int
}

func NewCK2Date(rawString string) *CK2Date {
	t, e := time.Parse("2006.1.2", rawString)
	checkError(e)
	return &CK2Date{t.Year(), t.Month(), t.Day()}
}
