package ck2save

import (
	"io"
)

func BuildSaveData(r io.Reader) *SaveData {
	s := NewSaveData()
	var cur, next *Property
	var ee error
	e := ReadSaveFile(r, func(l string) {
		next, ee = ParseProperty(l)
		if ee != nil {
			panic(ee)
		}
		if cur != nil {
			switch next.pattern {
			case newPropPattern, undefinedPattern:
			}
		} else if next.pattern == headerPattern {
			s.AddProperty(next)
			cur = next
		} else {
			panic(1000)
		}
	})
	if e != nil {
		if e != io.EOF {
			panic(e)
		}
	}
	return s
}
