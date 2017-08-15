package ck2save

import (
	"fmt"
)

type ck2SaveFileBuilder struct {
	curPropMap *propMap
	str        string
	save       *CK2Save
}

func newCK2SaveFileBuilder(s *CK2Save) ck2SaveFileBuilder {
	b := ck2SaveFileBuilder{}
	b.save = s
	return b
}

func (b *ck2SaveFileBuilder) buildSaveFile() string {
	b.curPropMap = b.save.propMapList[0]
	b.buildPropMap(b.curPropMap)
	return b.str
}

func (b *ck2SaveFileBuilder) buildPropMap(p *propMap) {
	switch p.pattern {
	case headerPattern:
		b.str += p.name + "\n"
	case newUnnamedMapPattern:
		b.str += "{\n"
	case newNamedMapPattern:
		b.str += p.name + "=\n"
	case newNamedMapSameLinePattern:
		b.str += p.name + "={\n"
	case undefinedPattern:
		b.buildPropMap(p.propMapList[0])
		return
	default:
		panic(b.str)
	}
	b.buildProperties(p)

	curPropMap := b.curPropMap
	b.curPropMap = p
	for _, pm := range b.curPropMap.propMapList {
		b.buildPropMap(pm)
	}

	switch p.pattern {
	case headerPattern, newUnnamedMapPattern, newNamedMapSameLinePattern:
		b.str += "}\n"
	case newNamedMapPattern, undefinedPattern:
	default:
		panic("Unexpected pattern for current propMap")
	}

	b.curPropMap = curPropMap
}

func (b *ck2SaveFileBuilder) buildProperties(p *propMap) {
	for _, pk := range p.propKeys {
		b.str += fmt.Sprintf("%v=%v\n", pk, p.property[pk])
	}
	//	for pn, pv := range p.property {
	//		b.str += pn + "=" + pv.(string) + "\n"
	//	}
}
