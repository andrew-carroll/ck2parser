package ck2save

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
	b.buildPropMap(&b.save.propMapList[0].propMapList[0])
	return b.str
}

func (b *ck2SaveFileBuilder) buildPropMap(p *propMap) {
	switch p.pattern {
	case headerPattern:
		b.str += "CK2txt\n"
	case newUnnamedMapPattern:
		b.str += "{\n"
	case newNamedMapPattern:
		b.str += p.name + "=\n"
	case newNamedMapSameLinePattern:
		b.str += p.name + "={\n"
	default:
		panic(b.str)
	}
	b.buildProperties(p)

	curPropMap := b.curPropMap
	b.curPropMap = p
	for _, pm := range b.curPropMap.propMapList {
		b.buildPropMap(&pm)
	}

	switch p.pattern {
	case headerPattern, newUnnamedMapPattern, newNamedMapSameLinePattern:
		b.str += "}\n"
	case newNamedMapPattern:
	default:
		panic("Unexpected pattern for current propMap")
	}

	b.curPropMap = curPropMap
}

func (b *ck2SaveFileBuilder) buildProperties(p *propMap) {
	for pn, pv := range p.property {
		b.str += pn + "=" + pv.(string) + "\n"
	}
}
