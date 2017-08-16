package ck2save

import (
	"bufio"
	"fmt"
)

type CK2SaveFileBuilder struct {
	curPropMap *propMap
	writer     *bufio.Writer
	save       *CK2Save
}

func NewCK2SaveFileBuilder(s *CK2Save) CK2SaveFileBuilder {
	b := CK2SaveFileBuilder{}
	b.save = s
	return b
}

func (b *CK2SaveFileBuilder) WriteSaveFile(filename string) {
	w, fClose := openFileWriter(filename)
	b.writer = w
	defer fClose()
	b.curPropMap = b.save.PropMapList[0]
	b.buildPropMap(b.curPropMap)
	b.writer.Flush()
}

func (b *CK2SaveFileBuilder) write(str string) {
	b.writer.WriteString(str)
}

func (b *CK2SaveFileBuilder) buildPropMap(p *propMap) {
	switch p.pattern {
	case headerPattern:
		b.write(p.name + "\n")
	case newUnnamedMapPattern:
		b.write("{\n")
	case newNamedMapPattern:
		b.write(p.name + "=\n")
	case newNamedMapSameLinePattern:
		b.write(p.name + "={\n")
	case undefinedPattern:
		b.buildPropMap(p.propMapList[0])
		return
	default:
		panic("Unhandled pattern")
	}
	b.buildProperties(p)

	curPropMap := b.curPropMap
	b.curPropMap = p
	for _, pm := range b.curPropMap.propMapList {
		b.buildPropMap(pm)
	}

	switch p.pattern {
	case headerPattern, newUnnamedMapPattern, newNamedMapSameLinePattern:
		b.write("}\n")
	case newNamedMapPattern, undefinedPattern:
	default:
		panic("Unexpected pattern for current propMap")
	}

	b.curPropMap = curPropMap
}

func (b *CK2SaveFileBuilder) buildProperties(p *propMap) {
	for _, pk := range p.propKeys {
		b.write(fmt.Sprintf("%v=%v\n", pk, p.property[pk]))
	}
	//	for pn, pv := range p.property {
	//		b.str += pn + "=" + pv.(string) + "\n"
	//	}
}
