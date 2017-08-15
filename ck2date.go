package ck2save

import (
	"fmt"
	"time"
)

type ck2Date struct {
	year   int
	month  int
	day    int
	quoted bool
}

func newCK2Date(rawString string, quoted bool) ck2Date {
	switch quoted {
	case true:
		timeFormat := `"2006.1.2"`
		t, e := time.Parse(timeFormat, rawString)
		checkError(e)
		return ck2Date{t.Year(), int(t.Month()), t.Day(), quoted}
	default:
		timeFormat := "2006.1.2"
		t, e := time.Parse(timeFormat, rawString)
		checkError(e)
		return ck2Date{t.Year(), int(t.Month()), t.Day(), quoted}
	}
}

func (d ck2Date) String() string {
	if d.quoted {
		return fmt.Sprintf("\"%d.%d.%d\"", d.year, d.month, d.day)
	} else {
		return fmt.Sprintf("%d.%d.%d", d.year, d.month, d.day)
	}
}
